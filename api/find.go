package api

import "github.com/train-cat/client-train-go"

var cacheStationName = map[int]string {}

// FindStationNameByID return station name
func FindStationNameByID(stationID int) string {
	name, exist := cacheStationName[stationID]

	if exist && name != "" {
		return name
	}

	s, err := traincat.GetStation(stationID)

	if err != nil {
		// TODO log
		return ""
	}

	cacheStationName[stationID] = s.Name

	return s.Name
}
