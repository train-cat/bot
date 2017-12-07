package api

import "github.com/train-cat/client-train-go"

// FindStationNameByID return station name
func FindStationNameByID(stationID int) string {
	s, err := traincat.GetStation(stationID)

	if err != nil {
		// TODO log
		return ""
	}

	return s.Name
}
