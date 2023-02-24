package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dayvillefire/newworld-cadview-agent/agent"
	"github.com/dayvillefire/newworld-cadview-utils/util"
)

var (
	configFile = flag.String("config", "config.yaml", "Configuration file")

	a              = map[string]*agent.Agent{}
	initialized    bool
	discordSession *discordgo.Session
	discordInit    bool
	stop           = make(chan os.Signal)
	shuttingDown   bool

	callMap        = map[int64]*discordgo.Channel{}
	lastUpdatedMap = map[int64]time.Time{}
	agentMap       = map[int64]string{}
)

type serFormat struct {
	CallMap        map[int64]discordgo.Channel
	LastUpdatedMap map[int64]time.Time
	AgentMap       map[int64]string
}

func main() {
	flag.Parse()

	var err error
	Config, err = LoadConfigWithDefaults(*configFile)
	if err != nil {
		panic(err)
	}

	log.Printf("INFO: Config: %#v", Config)

	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGQUIT)
	signal.Notify(stop, syscall.SIGINT)
	signal.Notify(stop, syscall.SIGHUP)

	err = initAgent()
	if err != nil {
		panic(err)
	}

	err = initDiscord(Config.Discord.Token)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(Config.Paths.SerializationFile)
	if err == nil {
		raw, err := ioutil.ReadFile(Config.Paths.SerializationFile)
		if err != nil {
			log.Printf("ERR: ReadFile(%s): %s", Config.Paths.SerializationFile, err.Error())
		}
		s, err := util.FromGOB64[serFormat](string(raw))
		//s, err := util.UnserializeFromFile[serFormat](*serialFile)
		if err != nil {
			log.Printf("ERR: FromG0B64(): %s", err.Error())
		}
		if err == nil {
			//log.Printf("TRACE: %#v", s)

			callMap = rerefChannelMap(s.CallMap)
			lastUpdatedMap = s.LastUpdatedMap
			agentMap = s.AgentMap

			log.Printf("INFO: Restored callMap : %#v", callMap)
			log.Printf("INFO: Restored lastUpdatedMap : %#v", lastUpdatedMap)
			log.Printf("INFO: Restored agentMap : %#v", agentMap)
		}
	}

	go func() {
		log.Printf("INFO[ReAuth]: Starting reauthorization thread")
		for {
			for i := 0; i < 60*Config.ReauthMinutes; i++ {
				time.Sleep(time.Second)
				if shuttingDown {
					log.Printf("INFO[ReAuth]: Shutting down")
					return
				}
			}
			for k, v := range a {
				reauth := v.MakeCopy()
				err := reauth.Init()
				if err != nil {
					log.Printf("ERR[ReAuth]: Unable to authenticate: %s", err.Error())
					continue
				}
				log.Printf("INFO[ReAuth]: Refreshing agent %s auth info", k)
				a[k].TransferAuthFrom(reauth)
			}

			if shuttingDown {
				log.Printf("INFO[ReAuth]: Shutting down")
				return
			}
		}
	}()

	// Polling
	go func() {
		for {
			active := []agent.CallObj{}
			for npa, pa := range a {
				log.Printf("INFO[%s]: Polling for active calls", npa)
				thisActive, err := pa.GetActiveCalls()
				if err != nil {
					log.Printf("ERR[%s]: GetActiveCalls(): %s", npa, err.Error())
					break
				}
				for _, ac := range thisActive {
					callAgent, found := agentMap[ac.CallID]
					if !found {
						// Add to agent map
						agentMap[ac.CallID] = npa
					}
					if callAgent == npa {
						// Skip if we have already added
						for _, y := range active {
							if y.CallID == ac.CallID {
								continue
							}
						}
						// Append to active list, otherwise
						active = append(active, ac)
					}
				}
			}
			for _, c := range active {
				log.Printf("INFO[%s]: Call = %d (%s - %s)", agentMap[c.CallID], c.CallID, c.Location, c.NatureOfCall)

				_, ok := callMap[c.CallID]
				if !ok {
					// Initial message ID
					m := discordgo.MessageSend{
						Content:         fmt.Sprintf("%s PRI %s (%s)", c.Location, c.CallPriority, c.CallType),
						AllowedMentions: &discordgo.MessageAllowedMentions{},
					}

					// Initial message, from which to build the "channel"
					res, err := discordSession.ChannelMessageSendComplex(Config.Discord.ChannelID, &m)
					if err != nil {
						log.Printf("ERR[%s]: ChannelMessageSendComplex(): %s", agentMap[c.CallID], err.Error())
						break
					}

					// Create complex message thread
					{
						t, err := discordSession.MessageThreadStartComplex(Config.Discord.ChannelID, res.ID, &discordgo.ThreadStart{
							Name: fmt.Sprintf("%s PRI %s (%s)", c.Location, c.CallPriority, c.CallType),
							//AutoArchiveDuration: 60,
							Invitable:        false,
							RateLimitPerUser: 10,
						})
						if err != nil {
							log.Printf("ERR[%s]: MessageThreadStartComplex(): %s", agentMap[c.CallID], err.Error())
						} else {
							callMap[c.CallID] = t
							lastUpdatedMap[c.CallID] = time.Now()
						}
					}

					// Send a message with initial times, etc
					{
						_, err = discordSession.ChannelMessageSendComplex(callMap[c.CallID].ID, &discordgo.MessageSend{
							Content:         fmt.Sprintf("Call dispatched at %s", c.DispatchedDateTime),
							AllowedMentions: &discordgo.MessageAllowedMentions{},
						})
						if err != nil {
							log.Printf("ERR[%s]: ChannelMessageSendComplex(%d): %s", agentMap[c.CallID], c.CallID, err.Error())
						}
					}
				} else {
					// Existing, so check for anything and append
					myAgent := a[agentMap[c.CallID]]
					log.Printf("INFO[%s]: Fetching call logs for call id %d", agentMap[c.CallID], c.CallID)
					logs, err := myAgent.GetCallLogs(fmt.Sprintf("%d", c.CallID))
					if err != nil {
						log.Printf("ERR[%s]: GetCallLogs(%d): %s", agentMap[c.CallID], c.CallID, err.Error())
						break
					}

					// Sort logs by date before doing this
					sort.SliceStable(logs, func(i, j int) bool {
						return logs[i].LogDateTime < logs[j].LogDateTime
					})

					for _, l := range logs {
						// LogDateTime:"02/10/2023 10:40:54"
						t, err := time.ParseInLocation("01/02/2006 15:04:05", l.LogDateTime, time.Local)
						//log.Printf("TRACE: t = %#v, l.LogDateTime = %s, lastUpdatedMap[c.CallID] = %#v", t, l.LogDateTime, lastUpdatedMap[c.CallID])
						//log.Printf("TRACE: l = %#v", l)
						if err != nil {
							log.Printf("WARN: Could not parse date %s", l.LogDateTime)
							continue
						}
						if t.After(lastUpdatedMap[c.CallID]) {
							found := false
							for _, st := range Config.LoggedStatuses {
								if st == l.ActionDescription {
									found = true
									_, err = discordSession.ChannelMessageSendComplex(callMap[c.CallID].ID, &discordgo.MessageSend{
										Content:         fmt.Sprintf("%s: %s (%s)", l.LogDateTime, l.Description, l.LastName),
										AllowedMentions: &discordgo.MessageAllowedMentions{},
									})
									if err != nil {
										log.Printf("ERR: ChannelMessageSendComplex(%d): %s", c.CallID, err.Error())
									}
								}
							}
							if !found {
								log.Printf("INFO: Skipping log item %#v", l)
							}

							// Update
							lastUpdatedMap[c.CallID] = t
						}
					}
					log.Printf("INFO[%s]: Call %d has %d log entries", agentMap[c.CallID], c.CallID, len(logs))
				}
			}

			log.Printf("INFO: Purging calls which have not been updated in over %d minutes", Config.PurgeMinutes)
			for k, v := range lastUpdatedMap {
				if time.Now().Local().Sub(v) > time.Minute*time.Duration(Config.PurgeMinutes) {
					log.Printf("INFO: Removing call %d with timestamp %d, now = %d", k, v.Unix(), time.Now().Local().Unix())
					delete(lastUpdatedMap, k)
					delete(callMap, k)
					delete(agentMap, k)
				}
			}

			if shuttingDown {
				log.Printf("INFO: Shutting down")
				return
			}

			time.Sleep(time.Duration(Config.PollInterval) * time.Second)
		}
	}()

	sig := <-stop
	log.Printf("INFO: Caught signal %#v", sig)

	toSerialize := serFormat{
		CallMap:        derefChannelMap(callMap),
		LastUpdatedMap: lastUpdatedMap,
		AgentMap:       agentMap,
	}

	log.Printf("INFO: Serializing %#v", toSerialize)
	data, err := util.ToGOB64[serFormat](toSerialize)
	if err != nil {
		log.Printf("ERR: ToGOB64(): %s", err.Error())
	}
	err = ioutil.WriteFile(Config.Paths.SerializationFile, []byte(data), 0644)
	if err != nil {
		log.Printf("ERR: WriteFile(%s): %s", Config.Paths.SerializationFile, err.Error())
	}

	os.Exit(0)
}

// Init initializes configured agent.
func initAgent() error {
	if initialized {
		return fmt.Errorf("already initialized")
	}

	for n, cd := range Config.CadInstances {
		log.Printf("INFO[Agent %s]: Initializing", n)
		thisAgent := agent.Agent{
			Debug:    false,
			LoginUrl: cd.URL,
			Username: cd.Username,
			Password: cd.Password,
			FDID:     cd.FDID,
		}
		err := thisAgent.Init()
		if err != nil {
			log.Printf("ERR[Agent %s]: Failed to intialized properly: %s", n, err)
		} else {
			a[n] = &thisAgent
		}
	}

	return nil
}

func initDiscord(token string) error {
	var err error
	if discordInit {
		return fmt.Errorf("ERR: already intiialized: %w", err)
	}

	discordSession, err = discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("ERR: discordgo.New(): %w", err)
	}

	err = discordSession.Open()
	if err != nil {
		return fmt.Errorf("ERR: discordSession.Open(): %w", err)
	}

	discordInit = true
	return nil
}

/*
func sendMessage(channel, msg string) (string, error) {
        m := discordgo.MessageSend{
                Content:         msg,
				Reference: ,
                AllowedMentions: &discordgo.MessageAllowedMentions{},
        }

        // Post normal message
        res, err := discordSession.ChannelMessageSendComplex(channel, &m)
        if err != nil {
                return "", err
        }

        return res.ID, nil
}
*/

func derefChannelMap(in map[int64]*discordgo.Channel) map[int64]discordgo.Channel {
	out := map[int64]discordgo.Channel{}
	for k, v := range in {
		out[k] = *v
	}
	return out
}

func rerefChannelMap(in map[int64]discordgo.Channel) map[int64]*discordgo.Channel {
	out := map[int64]*discordgo.Channel{}
	for k, v := range in {
		out[k] = &v
	}
	return out
}
