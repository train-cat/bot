package api

import (
	"time"

	"github.com/train-cat/bot/helper"
	"github.com/train-cat/client-train-go"
	"github.com/train-cat/client-train-go/filters"
)

const (
	maxStops          = 10
	precisionSchedule = 20
)

// SearchStops return list of stops
func SearchStops(originID int, destinationID int, schedule string) ([]traincat.Stop, error) {
	t, err := time.Parse("15:04:05", schedule)

	if err != nil {
		return nil, err
	}

	before := t.Add(time.Minute * precisionSchedule).Format("15:04")
	after := t.Add(time.Minute * precisionSchedule * -1).Format("15:04")

	f := &filters.Stop{
		Pagination: filters.Pagination{
			MaxPerPage: helper.Int(maxStops),
		},
		ScheduledBefore:       helper.String(before),
		ScheduledAfter:        helper.String(after),
		TrainThroughStationID: helper.Int(destinationID),
	}

	stops, err := traincat.CGetStops(uint(originID), f)

	if err != nil {
		return nil, err
	}

	return stops, nil
}

// SearchStations return list of stations (find by name)
func SearchStations(name string) ([]traincat.Station, error) {
	f := &filters.Station{
		Pagination: filters.Pagination{
			MaxPerPage: helper.Int(maxStops),
		},
		Name: helper.String(name),
	}

	stations, err := traincat.CGetAllStations(f)

	if err != nil {
		return nil, err
	}

	return stations, nil
}
