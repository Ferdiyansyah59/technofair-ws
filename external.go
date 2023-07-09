package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeatherResponse struct {
}

func CurrentWeather(lat float64, lon float64, appID string) (WeatherResponse, error) {
	requestURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", lat, lon, appID)
	res, err := http.Get(requestURL)

	if err != nil {
		fmt.Printf("error making http req: %s\n", err)
		return WeatherResponse{}, err
	}


	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Client: could not read response body %s\n", err)
		return WeatherResponse{}, err
	}

	data := WeatherResponse{}
	err = json.Unmarshal(resBody, &data)
	if err != nil {
		fmt.Printf("Json unmarshal error %s\n", err)
		return WeatherResponse{}, err
	}
	
	return data, nil
}