package models

type Train struct {
	Number        string `json:"number"`
	Title         string `json:"title"`
	StationFrom   string `json:"station_from"`
	DepartureTime string `json:"departure_time,omitempty"`
	ArrivalTime   string `json:"arrival_time,omitempty"`
	StationTo     string `json:"station_to,omitempty"`
}

type Station struct {
	Title     string  `json:"title"`
	Code      string  `json:"code"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
