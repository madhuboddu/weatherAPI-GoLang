package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func GetWeather(zipcode string) string {

	responseBody := GetWeatherInfo(zipcode)
	var fields = JsonMarshal(responseBody)
	formatMessgae := DisplayMessage(fields)

	return formatMessgae
}

func GetWeatherInfo(zipcode string) []byte {

	config, err := LoadConfig(".")

	if err != nil {
		return nil
	}

	baseURL := config.BaseURL
	apiKey := config.APIKey
	location := zipcode
	aqi := "yes"

	encodedLocation := url.QueryEscape(location)

	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=%s", baseURL, apiKey, encodedLocation, aqi)
	fmt.Print(url)
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := client.Do(request)
	defer response.Body.Close()
	responseBody, _ := io.ReadAll(response.Body)

	return responseBody
}

func home(c *gin.Context) {
	// Popup.Alert("Weather App", "This is a simple weather app. The app gives you the temperature in Celcius at the give Canadian Postal code. Try it your self. ex : N2J")
	fmt.Print("This is a simple weather app. The app gives you the temperature in Celcius at the give Canadian Postal code. Try it your self. ex : N2J")
}

func GetWeatherByCode(c *gin.Context) string {
	zipcode := c.Param("zipcode")

	responseBody := GetWeatherInfo(zipcode)
	var fields = JsonMarshal(responseBody)
	formatMessgae := DisplayMessage(fields)

	c.IndentedJSON(http.StatusOK, formatMessgae)

	if formatMessgae != "" {
		c.IndentedJSON(http.StatusOK, formatMessgae)
		return "Sucess : Request OK"
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ZIPCODE not found"})

	return formatMessgae
}
