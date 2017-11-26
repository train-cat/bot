package api

import (
	"github.com/train-cat/bot/helper"
	"github.com/train-cat/client-train-go"
	"github.com/train-cat/client-train-go/filters"
)

// FindStationNameByID return station name
func FindStationNameByID(stationID int) string {
	s, err := traincat.GetStation(stationID)

	if err != nil {
		// TODO log
		return ""
	}

	return s.Name
}

// FindStopByOriginAndSchedule return one stop
func FindStopByOriginAndSchedule(originID int, schedule string) *traincat.Stop {
	// TODO maybe add destination
	f := &filters.Stop{
		Pagination: filters.Pagination{
			MaxPerPage: helper.Int(1),
		},
		ScheduledAt: helper.String(schedule),
	}

	stops, err := traincat.CGetStops(uint(originID), f)

	if err != nil {
		// TODO log
		return nil
	}

	if len(stops) == 0 {
		return nil
	}

	return &stops[0]
}
