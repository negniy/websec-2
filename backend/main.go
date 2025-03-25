package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const apiURL = "https://api.rasp.yandex.net/v3.0/search/"

func getTrainSchedule() (map[string]interface{}, error) {
	// Параметры запроса
	params := url.Values{}
	params.Add("apikey", "4d138079-8d8f-4e5f-8b60-8471062e5365")
	params.Add("from", "c146")
	params.Add("to", "c213")
	params.Add("date", "2025-09-02")

	// Формируем URL с параметрами
	fullURL := apiURL + "?" + params.Encode()

	// Отправляем GET-запрос
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе: %v", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: получен статус %d", resp.StatusCode)
	}

	// Парсим JSON-ответ
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("ошибка при разборе JSON: %v", err)
	}

	return result, nil
}

func main() {
	// Получаем расписание
	scheduleData, err := getTrainSchedule()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	// Выводим данные
	fmt.Printf("Расписание: %+v\n", scheduleData)
}
