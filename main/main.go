package main

import (
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
	router.GET("/weather", home)
	router.Run("localhost:8080")
}
