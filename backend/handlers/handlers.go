package handlers

import (
	api "backend/API"
	"encoding/json"
	"log"
	"net/http"
)

// Обработчик запроса для получения всех поездов через станцию
func FetchTrainsThroughStation(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из запроса
	station := r.URL.Query().Get("station")
	date := r.URL.Query().Get("date")

	// Вызываем функцию из API для получения данных
	trains, err := api.GetTrainsThroughStation(api.Stations[station], date)
	if err != nil {
		// Возвращаем ошибку сервера
		http.Error(w, "Ошибка при получении данных о поездах", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trains) // Отправляем данные в формате JSON
}

// Обработчик запроса для получения расписания поездов между двумя станциями
func FetchTrainRoute(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из запроса
	fromStation := r.URL.Query().Get("from")
	toStation := r.URL.Query().Get("to")
	date := r.URL.Query().Get("date")

	// Вызываем функцию из API для получения данных
	log.Println(api.Stations[fromStation], api.Stations[toStation])
	trains, err := api.GetTrainRoute(api.Stations[fromStation], api.Stations[toStation], date)
	if err != nil {
		// Возвращаем ошибку сервера
		http.Error(w, "Ошибка при получении расписания", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовки для ответа
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(trains) // Отправляем данные в формате JSON
}
