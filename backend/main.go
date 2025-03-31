package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "backend/API"
	"backend/handlers"

	"github.com/gorilla/mux"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}

func init() {
	api.LoadStation()
	log.Println(api.Stations)
}

func main() {
	router := mux.NewRouter()

	router.Use(enableCORS)

	router.HandleFunc("/api/trains/through/", handlers.FetchTrainsThroughStation)
	router.HandleFunc("/api/trains/route/", handlers.FetchTrainRoute)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Сервер слушает на порту 8080...")
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}
