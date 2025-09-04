package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dayvillefire/newworld-cadview-agent/agent"
	"github.com/jbuchbinder/shims"
)

func stats() {
	cc, err := readcalls()
	if err != nil {
		log.Printf("ERR: %s", err.Error())
		return
	}

	fmt.Printf("Current data contains %d calls\n", len(cc))

	// Determine stop/start point
	{
		min := "X"
		max := ""
		shims.ArrayFunc(cc, func(in agent.CADCall) {
			//log.Printf("%#v", in.Call)
			if in.Call.CreatedDateTime < min {
				min = in.Call.CreatedDateTime
			}
			if in.Call.CreatedDateTime > max {
				max = in.Call.CreatedDateTime
			}
		})
		fmt.Printf("Range: %s to %s\n\n", min, max)
	}

	{
		missedcalls := 0
		maonly := 0
		missedoffhourcalls := 0
		totalcalls := len(cc)
		totaloffhourcalls := 0
		unitcount := map[string]int{}

		shims.ArrayFunc(cc, func(in agent.CADCall) {
			offhours := false
			incidentNumber := ""
			shims.ArrayFunc(in.Incidents, func(inc agent.IncidentObj) {
				if inc.ORI == *ori {
					incidentNumber = inc.IncidentNumber
				}
			})

			// Determine if this is off hours
			dt, err := time.Parse("01/02/2006 15:04:05", in.Call.CreatedDateTime)
			if err == nil {
				// Look for day of the week
				if dt.Hour() >= 8 || dt.Hour() < 16 {
					if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
						offhours = true
						totaloffhourcalls += 1
					}
				}
			}

			oh := ""
			if offhours {
				oh = " [OFFHOURS]"
			}
			fmt.Printf("Call: %s @ %s (%s)%s\n", incidentNumber, in.Call.Location, in.Call.CreatedDateTime, oh)
			fmt.Printf("\tPriority: %s | %s | %s\n", in.Call.CallPriority, in.Call.FireCallType, in.Call.Quadrant)

			un := []string{}
			maun := []string{}
			enroute := ""
			onscene := ""
			shims.ArrayFunc(in.Units, func(u agent.UnitObj) {
				for _, p := range strings.Split(*skipprefix, ",") {
					if strings.HasPrefix(u.UnitNumber, strings.TrimSpace(p)) {
						return
					}
				}
				if *masuffix != "" && strings.HasSuffix(u.UnitNumber, strings.TrimSpace(*masuffix)) {
					maun = append(maun, u.UnitNumber)
					return
				}
				if *reqsuffix != "" && !strings.HasSuffix(u.UnitNumber, strings.TrimSpace(*reqsuffix)) {
					return
				}
				un = append(un, u.UnitNumber)
				{
					unitcount[u.UnitNumber] += 1
				}

				// Determine if enroute / onscene times are earliest
				if u.EnrouteDateTime != "" {
					if enroute == "" || enroute < u.EnrouteDateTime {
						enroute = u.EnrouteDateTime
					}
				}
				if u.ArriveDateTime != "" {
					if onscene == "" || onscene < u.ArriveDateTime {
						onscene = u.ArriveDateTime
					}
				}
			})
			if len(un) > 0 {
				fmt.Printf("\tUnits: %s\n", strings.Join(un, ","))
				if enroute != "" {
					er, err := time.Parse("01/02/2006 15:04:05", enroute)
					if err == nil {
						fmt.Printf("\tEnroute time : %s\n", er.Sub(dt))
					}
				}
				if onscene != "" {
					ons, err := time.Parse("01/02/2006 15:04:05", onscene)
					if err == nil {
						fmt.Printf("\tOn scene time : %s\n", ons.Sub(dt))
					}
				} else {
					fmt.Printf("\tCANCELLED EN ROUTE\n")
				}
			} else {
				fmt.Printf("\tMISSED CALL\n")
				missedcalls += 1
				if offhours {
					missedoffhourcalls += 1
				}
				if len(maun) > 0 {
					maonly += 1
				}
			}
			if len(maun) > 0 {
				fmt.Printf("\tMutual Aid Units: %s\n", strings.Join(maun, ", "))
			}
			fmt.Println("")
		})
		fmt.Printf("\n")
		fmt.Printf("Aggregate => Total calls: %d, Missed calls: %d, Missed call %%: %.2f\n",
			totalcalls, missedcalls,
			float64(float64(missedcalls)/float64(totalcalls))*100)
		fmt.Printf("Off Hours => Total calls: %d, Missed calls: %d, Missed call %%: %.2f\n",
			totaloffhourcalls, missedoffhourcalls,
			float64(float64(missedoffhourcalls)/float64(totaloffhourcalls))*100)
		if *masuffix != "" {
			fmt.Printf("Mutual Aid Only: %d calls\n", maonly)
		}

		for unit, count := range unitcount {
			fmt.Printf("\t%s : %d\n", unit, count)
		}
	}
}
