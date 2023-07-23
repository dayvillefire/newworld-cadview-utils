package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	channelId   = flag.String("channel", "1074021838088323112", "Channel ID")
	matchText   = flag.String("matchText", "(Service Call)", "Text to match to delete messages. Leave blank to delete all.")
	threadPurge = flag.Bool("threadPurge", false, "Also purge threads explicitly")

	initialized    bool
	discordSession *discordgo.Session
	discordInit    bool
)

func main() {
	flag.Parse()

	t := true

	err := initDiscord(token)
	if err != nil {
		panic(err)
	}

	for {
		msgs, err := discordSession.ChannelMessages(*channelId, 100, "", "", "")
		if err != nil {
			panic(err)
		}

		log.Printf("Found %d channel messages (max batch 100)", len(msgs))

		if len(msgs) < 1 {
			break
		}

		del := make([]string, 0)

		for k, v := range msgs {
			log.Printf("Found message #%d, text %s, ID %s", k, v.Content, v.ID)
			if v.Thread != nil {
				if v.Thread.MessageCount > 2 {
					log.Printf("Starts thread, skipping")
					continue
				}
			}
			if *matchText != "" && strings.Index(v.Content, *matchText) == -1 {
				continue
			}

			// Add to list of messages to bulk delete
			del = append(del, v.ID)
		}

		if len(del) == 0 {
			break
		}

		log.Printf("Purging messages #%v", del)
		err = discordSession.ChannelMessagesBulkDelete(*channelId, del)
		if err != nil {
			panic(err)
		}

	}

	if !*threadPurge {
		log.Printf("Set not to purge threads explicitly, exiting")
		return
	}

	for {
		threads, err := discordSession.ThreadsActive(*channelId)
		if err != nil {
			panic(err)
		}

		log.Printf("Threads: %d found", len(threads.Threads))
		for k, v := range threads.Threads {
			log.Printf("[%d] %s (%d messages)", k, v.Topic, v.MessageCount)
			if v.MessageCount > 2 {
				continue
			}
			_, err := discordSession.ChannelEdit(v.ID, &discordgo.ChannelEdit{Archived: &t})
			if err != nil {
				log.Printf("ERR: %s", err.Error())
			}
		}

		if !threads.HasMore {
			log.Printf("No more threads, exiting")
			break
		}

		time.Sleep(time.Second)

		log.Printf("More threads, continuing")
	}

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
