package util

import (
	"strings"
	"time"

	"github.com/dayvillefire/newworld-cadview-agent/agent"
	"github.com/jbuchbinder/shims"
)

const (
	CadTimeFormat = "01/02/2006 15:04:05"
)

type CallDataProcessor struct {
	Call          agent.CADCall
	SkipPrefix    []string
	RequireSuffix string
}

func (c CallDataProcessor) DispatchedDateTime(agency string) (time.Time, error) {
	dispatched := ""

	shims.ArrayFunc(c.Call.Units, func(u agent.UnitObj) {
		if u.UnitNumber != "RES"+agency && u.UnitNumber != "STA"+agency {
			return
		}

		// Determine if enroute / onscene times are earliest
		if u.DispatchDateTime != "" {
			if dispatched == "" || u.DispatchDateTime < dispatched {
				dispatched = u.DispatchDateTime
			}
		}
	})
	return time.Parse(CadTimeFormat, dispatched)
}

func (c CallDataProcessor) IsOffHours(agency string) bool {
	dt, err := c.DispatchedDateTime(agency)
	if err == nil {
		// Look for day of the week
		if dt.Hour() >= 8 || dt.Hour() < 16 {
			if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
				return true
			}
		}
	}
	return false
}

func (c CallDataProcessor) UnitIdentifiers(agency string) []string {
	un := []string{}
	shims.ArrayFunc(c.Call.Units, func(u agent.UnitObj) {
		for _, p := range c.SkipPrefix {
			if strings.HasPrefix(u.UnitNumber, strings.TrimSpace(p)) {
				return
			}
		}
		if c.RequireSuffix != "" && !strings.HasSuffix(u.UnitNumber, strings.TrimSpace(c.RequireSuffix)) {
			return
		}
		un = append(un, u.UnitNumber)
	})
	return un
}

func (c CallDataProcessor) EnrouteDateTime(agency string) (time.Time, error) {
	un := []string{}
	enroute := ""
	shims.ArrayFunc(c.Call.Units, func(u agent.UnitObj) {
		for _, p := range c.SkipPrefix {
			if strings.HasPrefix(u.UnitNumber, strings.TrimSpace(p)) {
				return
			}
		}
		if c.RequireSuffix != "" && !strings.HasSuffix(u.UnitNumber, strings.TrimSpace(c.RequireSuffix)) {
			return
		}
		un = append(un, u.UnitNumber)

		// Determine if enroute / onscene times are earliest
		if u.EnrouteDateTime != "" {
			if enroute == "" || enroute < u.EnrouteDateTime {
				enroute = u.EnrouteDateTime
			}
		}
	})
	return time.Parse(CadTimeFormat, enroute)
}

func (c CallDataProcessor) ArrivalDateTime(agency string) (time.Time, error) {
	un := []string{}
	onscene := ""
	shims.ArrayFunc(c.Call.Units, func(u agent.UnitObj) {
		for _, p := range c.SkipPrefix {
			if strings.HasPrefix(u.UnitNumber, strings.TrimSpace(p)) {
				return
			}
		}
		if c.RequireSuffix != "" && !strings.HasSuffix(u.UnitNumber, strings.TrimSpace(c.RequireSuffix)) {
			return
		}
		un = append(un, u.UnitNumber)

		// Determine if enroute / onscene times are earliest
		if u.ArriveDateTime != "" {
			if onscene == "" || onscene < u.ArriveDateTime {
				onscene = u.ArriveDateTime
			}
		}
	})
	return time.Parse(CadTimeFormat, onscene)
}
