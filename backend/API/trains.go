package api

import (
	"backend/models"
	"fmt"
)

type apiResponseThrough struct {
	Station  stationOne `json:"station"`
	Schedule []schedule `json:"schedule"`
}

type stationOne struct {
	Title string `json:"title"`
}

type direction struct {
	Title string `json:"title"`
}

type schedule struct {
	Arrival   string    `json:"arrival"`
	Thread    threadOne `json:"thread"`
	Departure string    `json:"departure"`
	Direction direction `json:"schedule_direction"`
}

type threadOne struct {
	Title  string `json:"title"`
	Number string `json:"number"`
}

func GetTrainsThroughStation(stationCode string, date string) ([]models.Train, error) {
	apiResponse, err := getTrainsThrough(stationCode, date)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения поездов через станцию: %v", err)
	}

	var trains []models.Train

	for _, schedule := range apiResponse.Schedule {
		train := models.Train{
			Number:        schedule.Thread.Number,
			Title:         schedule.Thread.Title,
			StationFrom:   apiResponse.Station.Title,
			DepartureTime: schedule.Departure,
			ArrivalTime:   schedule.Arrival,
			StationTo:     schedule.Direction.Title,
		}

		trains = append(trains, train)

	}

	return trains, nil
}

type apiResponseRoute struct {
	Segments []segment `json:"segments"`
}

type segment struct {
	Arrival   string      `json:"arrival"`
	From      stationInfo `json:"from"`
	Departure string      `json:"departure"`
	To        stationInfo `json:"to"`
	Thread    threadTwo   `json:"thread"`
}

type threadTwo struct {
	Number string `json:"number"`
	Title  string `json:"title"`
}

type stationInfo struct {
	Title string `json:"title"`
}

func GetTrainRoute(fromStationCode string, toStationCode string, date string) ([]models.Train, error) {
	apiResponse, err := getTrainsRoute(fromStationCode, toStationCode, date)
	if err != nil {
		return nil, fmt.Errorf("error getting trains through station: %v", err)
	}
	//log.Println(apiResponse)
	var trains []models.Train

	for _, segment := range apiResponse.Segments {
		train := models.Train{
			Number:        segment.Thread.Number,
			Title:         segment.Thread.Title,
			StationFrom:   segment.From.Title,
			DepartureTime: segment.Departure,
			ArrivalTime:   segment.Arrival,
			StationTo:     segment.To.Title,
		}

		//log.Println(train)
		trains = append(trains, train)

	}

	return trains, nil
}
