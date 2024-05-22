package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getWeatherByCode(c *gin.Context) {
	zipcode := c.Param("zipcode")

	responseBody := GetWeatherInfo(zipcode)

	fmt.Println(string(responseBody))
	var fields = JsonMarshal(responseBody)
	formatMessgae := DisplayMessage(fields)

	c.IndentedJSON(http.StatusOK, formatMessgae)

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	if formatMessgae != "" {
		c.IndentedJSON(http.StatusOK, formatMessgae)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "ZIPCODE not found"})
}
