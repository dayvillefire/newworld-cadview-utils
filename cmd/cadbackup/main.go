package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dayvillefire/newworld-cadview-agent/agent"
)

var (
	a           agent.Agent
	initialized bool

	url       = flag.String("url", "https://cadview.qvec.org/NewWorld.CadView", "CAD login URL")
	user      = flag.String("username", "", "CAD username")
	pass      = flag.String("password", "", "CAD password")
	fdid      = flag.String("fdid", "", "FDID")
	startDate = flag.String("start", "", "Start date in YYYY-MM-DD")
	endDate   = flag.String("end", "", "End date in YYYY-MM-DD (ends at 00:00 on that date)")
	basedir   = flag.String("basedir", "store", "Base directory for storage")
)

func main() {
	flag.Parse()

	if *user == "" || *pass == "" || *fdid == "" || *startDate == "" || *endDate == "" || *basedir == "" {
		flag.PrintDefaults()
		return
	}

	startDt, err := time.Parse("2006-01-02", *startDate)
	if err != nil {
		panic(err)
	}
	endDt, err := time.Parse("2006-01-02", *endDate)
	if err != nil {
		panic(err)
	}

	dates := getDateList(startDt, endDt)

	err = initAgent()
	if err != nil {
		panic(err)
	}

	oris, err := a.GetORIs()
	ori := agent.FDIDToORI(oris, *fdid)

	for _, dts := range dates {
		st, _ := time.Parse("2006-01-02", dts[0])
		et, _ := time.Parse("2006-01-02", dts[1])
		co, err := a.GetClearedCalls(st, et, ori)
		if err != nil {
			log.Printf("ERR: %v - %v: %s", st, et, err.Error())
			continue
		}
		for _, c := range co {
			year := dts[0][:4]
			cadcall, err := a.RetrieveCADCall(c)
			if err != nil {
				log.Printf("ERR: %s", err.Error())
				continue
			}
			var iid string
			for _, ins := range cadcall.Incidents {
				if ins.ORI == *fdid {
					iid = strings.TrimSpace(ins.IncidentNumber)
				}
			}

			b := *basedir + string(os.PathSeparator) + *fdid + string(os.PathSeparator) + year
			os.MkdirAll(b, 0755)

			var data []byte
			data, err = json.Marshal(cadcall)
			if err != nil {
				log.Printf("ERR: %s", err.Error())
				continue
			}

			log.Printf("INFO: Writing call %s to %s/%s (%d bytes)", iid, b, iid, len(data))
			err = ioutil.WriteFile(b+string(os.PathSeparator)+iid, data, 0644)
			if err != nil {
				log.Printf("ERR: %s", err.Error())
				continue
			}
		}
	}

}

func getDateList(st, et time.Time) [][]string {
	out := make([][]string, 0)
	nt := st
	for {
		if nt.Equal(et) || nt.After(et) {
			return out
		}

		out = append(out, []string{nt.Format("2006-01-02"), nt.Add(time.Hour * 24).Format("2006-01-02")})

		nt = nt.Add(time.Hour * 24)
	}
}

// Init initializes configured agent.
func initAgent() error {
	if initialized {
		return fmt.Errorf("already initialized")
	}

	log.Printf("INFO: Initializing agent")
	a = agent.Agent{
		Debug:    false,
		LoginUrl: *url,
		Username: *user,
		Password: *pass,
		FDID:     *fdid,
	}
	err := a.Init()
	if err != nil {
		log.Printf("ERR: Failed to intialized properly: %s", err)
		return err
	}

	return nil
}
