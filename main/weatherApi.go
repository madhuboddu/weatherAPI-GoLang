package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func GetWeatherInfo(zipcode string) []byte {

	baseURL := "http://api.weatherapi.com/v1/current.json"
	apiKey := "a489e9b203534843ba000645242604"
	location := zipcode
	aqi := "yes"

	encodedLocation := url.QueryEscape(location)

	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=%s", baseURL, apiKey, encodedLocation, aqi)

	println(url)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := client.Do(request)
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)

	return responseBody
}
