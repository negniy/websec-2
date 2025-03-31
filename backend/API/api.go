package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const apiURL = "https://api.rasp.yandex.net/v3.0/"
const apiKey = "4d138079-8d8f-4e5f-8b60-8471062e5365"
const defaultLat = 53.196013
const defaultLng = 50.099892
const distance = 50

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getTrainsThrough(stationCode string, date string) (*apiResponseThrough, error) {

	params := url.Values{}
	params.Add("station", stationCode)
	params.Add("date", date)
	params.Add("transport_types", "suburban")
	curApiURL := apiURL + "schedule/?" + params.Encode()

	req, err := http.NewRequest("GET", curApiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("Authorization", apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: получен статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	var result apiResponseThrough
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("ошибка при разборе JSON: %v", err)
	}

	return &result, nil
}

func getTrainsRoute(fromStationCode string, toStationCode string, date string) (*apiResponseRoute, error) {
	params := url.Values{}
	params.Add("from", fromStationCode)
	params.Add("to", toStationCode)
	params.Add("date", date)
	params.Add("transport_types", "suburban")
	curApiURL := apiURL + "search/?" + params.Encode()

	req, err := http.NewRequest("GET", curApiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("Authorization", apiKey)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: получен статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	var result apiResponseRoute
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("ошибка при разборе JSON: %v", err)
	}

	return &result, nil
}
