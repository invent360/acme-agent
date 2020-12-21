package models

import "time"

type APIObservation struct {
	Start     time.Time                `json:"start"` // Start Time of observation
	End       time.Time                `json:"end"`   // End Time of observation
	Endpoints []APIObservationEndpoint `json:"endpoints"`
}