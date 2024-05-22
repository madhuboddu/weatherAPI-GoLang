package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Current struct {
	Temperature float64 `json:"temp_c"`
	FeelsLike   float64 `json:"feelslike_c"`
	UVIndex     float64 `json:"uv"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Fields struct {
	Current  Current  `json:"current"`
	Location Location `json:"location"`
}

func main() {

	router := gin.Default()
	router.GET("/weather/:zipcode", getWeatherByCode)
	router.Run("localhost:8080")

	// fmt.Print("################ ! Welcome to Advanced weather application ! ################### \n")

	// var zipcode string
	// fmt.Print("Please enter zip code to view the Report\n")
	// fmt.Scan(&zipcode)

	// responseBody := GetWeatherInfo(zipcode)

	// fmt.Println(string(responseBody))
	// var fields = JsonMarshal(responseBody)
	// DisplayMessage(fields)
}

func DisplayMessage(fields Fields) string {
	var country = fields.Location.Country
	var location = fields.Location.Name
	var temperature = fields.Current.Temperature
	var feelsLike = fields.Current.FeelsLike

	// fmt.Printf("The temperature at %s - %s is %f but feels like %f\n", location, country, temperature, feelsLike)
	formatted := fmt.Sprintf("The temperature at %s - %s is %f but feels like %f\n", location, country, temperature, feelsLike)
	return formatted
}

func JsonMarshal(res []byte) Fields {

	var fields Fields
	err := json.Unmarshal(res, &fields)

	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	return fields
}
