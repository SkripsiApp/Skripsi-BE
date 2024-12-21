package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type RajaOngkirResponse struct {
	Rajaongkir struct {
		Results []struct {
			CityID   string `json:"city_id"`
			CityName string `json:"city_name"`
		} `json:"results"`
	} `json:"rajaongkir"`
}

func GetCityID(cityName string) (string, error) {
	url := "https://api.rajaongkir.com/starter/city"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("RAJAONGKIR_API_KEY")
	req.Header.Add("key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var rajaOngkirResp RajaOngkirResponse
	err = json.NewDecoder(resp.Body).Decode(&rajaOngkirResp)
	if err != nil {
		return "", err
	}

	for _, city := range rajaOngkirResp.Rajaongkir.Results {
		if city.CityName == cityName {
			return city.CityID, nil
		}
	}

	return "", fmt.Errorf("city not found")
}