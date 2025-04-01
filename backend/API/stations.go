package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var Stations map[string]string = make(map[string]string)

type apiStationsResponse struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	Regions []Region `json:"regions"`
	Title   string   `json:"title"`
}

type Region struct {
	Title       string       `json:"title"`
	Settlements []Settlement `json:"settlements"`
}

type Settlement struct {
	Stations []stationWide `json:"stations"`
}

type stationWide struct {
	Title         string `json:"title"`
	TransportType string `json:"transport_type"`
	Codes         struct {
		Code string `json:"yandex_code"`
	} `json:"codes"`
}

func LoadStation() {
	URL := fmt.Sprintf("%sstations_list/?apikey=%s", apiURL, apiKey)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("Ошибка при создании запроса: %v", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Ошибка при отправке запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка: получен статус %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	var result apiStationsResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Ошибка при разборе JSON: %v", err)
	}

	for _, country := range result.Countries {
		if strings.Compare(country.Title, "Россия") == 0 {
			for _, region := range country.Regions {
				if strings.Compare(region.Title, "Самарская область") == 0 {
					for _, settelment := range region.Settlements {
						for _, station := range settelment.Stations {
							if strings.Compare(station.TransportType, "train") == 0 {
								Stations[station.Title] = station.Codes.Code
							}
						}
					}
				}
			}
		}
	}
}
