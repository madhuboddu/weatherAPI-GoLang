package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonMarshal(res []byte) Fields {

	var fields Fields
	err := json.Unmarshal(res, &fields)

	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err)
	}

	return fields
}

func DisplayMessage(fields Fields) string {
	var country = fields.Location.Country
	var location = fields.Location.Name
	var temperature = fields.Current.Temperature
	var feelsLike = fields.Current.FeelsLike

	formatted := fmt.Sprintf("The temperature at %s - %s is %f but feels like %f\n", location, country, temperature, feelsLike)
	return formatted
}
