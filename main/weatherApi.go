package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetWeatherInfos() string {

	url := "http://api.weatherapi.com/v1/current.json?key=a489e9b203534843ba000645242604&q=N2J 2C1&aqi=yes"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, _ := http.NewRequest(http.MethodGet, url, nil)

	response, _ := client.Do(request)

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	fmt.Printf(string(body))

	return string(body)
}
