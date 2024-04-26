package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
	"net/url"
	"time"
)

type current struct {
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feelsLike"`
	UVIndex     int32   `json:"uv"`
}

type location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Fields struct {
	current  current
	location location
}

func main() {

	fmt.Print("################ ! Welcome to Advanced weather application ! ################### \n")
	responseBody := GetWeatherInfo()

	fmt.Println(string(responseBody))
	var fields = JsonMarshal(responseBody)
	DisplayMessage(fields)
}

func DisplayMessage(fields Fields) {
	var country = fields.location.Country
	var location = fields.location.Name
	var temperature = fields.current.Temperature
	var feelsLike = fields.current.FeelsLike

	fmt.Printf("The temperature at %s - %s is %f but feels like %f\n", location, country, temperature, feelsLike)

}

func JsonMarshal(res []byte) Fields {

	var fields Fields
	err := json.Unmarshal(res, &fields)

	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	return fields
}

func GetWeatherInfo() []byte {

	baseURL := "http://api.weatherapi.com/v1/current.json"
	apiKey := "a489e9b203534843ba000645242604"
	location := "N2J 2C1"
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
