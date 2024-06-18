package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
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

	myApp := app.New()
	w := myApp.NewWindow("Weather Application")
	w.Resize(fyne.NewSize(300, 400))

	background := canvas.NewImageFromFile("image.png")
	background.FillMode = canvas.ImageFillStretch

	hello := widget.NewLabel("You are in the most advanced weather report application")
	hello.Resize(fyne.NewSize(50, 100))
	zipcodeEntry := widget.NewEntry()
	zipcodeEntry.SetPlaceHolder("Enter a Canadian Zip Code")

	getWeatherButton := widget.NewButton("Get Weather", func() {
		zipcode := zipcodeEntry.Text
		weatherData := getWeather(zipcode)
		hello.SetText(weatherData)
	})

	content := container.NewMax(
		background, container.NewVBox(
			hello,
			zipcodeEntry,
			getWeatherButton,
		),
	)

	w.SetContent(content)
	w.ShowAndRun()

}
