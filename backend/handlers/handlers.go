package handlers

import (
	api "backend/API"
	"backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func FetchTrainsThroughStation(w http.ResponseWriter, r *http.Request) {
	station := r.URL.Query().Get("station")
	date := r.URL.Query().Get("date")

	trains, err := api.GetTrainsThroughStation(api.Stations[station], date)
	if err != nil {
		http.Error(w, "Ошибка при получении данных о поездах", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trains)
}

func FetchTrainRoute(w http.ResponseWriter, r *http.Request) {
	fromStation := r.URL.Query().Get("from")
	toStation := r.URL.Query().Get("to")
	date := r.URL.Query().Get("date")

	log.Println(api.Stations[fromStation], api.Stations[toStation])
	trains, err := api.GetTrainRoute(api.Stations[fromStation], api.Stations[toStation], date)
	if err != nil {
		http.Error(w, "Ошибка при получении расписания", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trains)
}

func GetStations(w http.ResponseWriter, r *http.Request) {

	log.Println("get stations")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	stations := make([]models.Station, 0, 10)
	for key, val := range api.Stations {
		stations = append(stations, models.Station{
			Title: key,
			Code:  val,
		})
	}
	json.NewEncoder(w).Encode(stations)
}
