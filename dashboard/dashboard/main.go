package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	owm "github.com/briandowns/openweathermap"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

// kindle 4 horizontal screen size
const width = 800
const height = 600

const margin = 20

const maxWidth = width - margin
const maxHeight = height - margin

const fontPath = "../Fonts/Lato-Regular.ttf"

type Config struct {
	Lat               float64 `json:"lat"`
	Lon               float64 `json:"lon"`
	OpenweatherApiKey string  `json:"openweatherApiKey"`
	// HistoryToday      HistoryToday `json:"historyToday"`
}

// stores the current weather data
var currentWeather *owm.CurrentWeatherData = &owm.CurrentWeatherData{}
var lastWeatherCheck time.Time = time.Time{}
var coords *owm.Coordinates = &owm.Coordinates{}
var todaysHistory HistoryToday

func main() {

	configFile, err := ioutil.ReadFile("../config.json")
	if err != nil {
		panic(err)
	}

	config := Config{}
	_ = json.Unmarshal(configFile, &config)

	coords = &owm.Coordinates{
		Latitude:  config.Lat,
		Longitude: config.Lon,
	}

	loadWeather(config.OpenweatherApiKey)

	todaysHistory, err = GetHistory()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			sleepTime := (60 - time.Now().Minute()) % 30
			// checks for weather every half an hour
			if time.Since(lastWeatherCheck).Minutes() >= 30 {
				currentWeather.CurrentByCoordinates(coords)
				lastWeatherCheck = time.Now()
			}
			time.Sleep(time.Minute * time.Duration(sleepTime))
		}

	}()

	go func() {
		for {
			// sleep until the next 24 hour time
			sleepTime := (24 - time.Now().Hour()) % 24
			time.Sleep(time.Hour * time.Duration(sleepTime))

			todaysHistory, _ = GetHistory()
		}
	}()

	for {

		dc := gg.NewContext(width, height)
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		dc.SetRGB(0, 0, 0)
		if err := dc.LoadFontFace(fontPath, 20); err != nil {
			panic(err)
		}

		drawDate(dc)
		drawHistory(dc, todaysHistory)

		drawWeather(dc)

		// draw last to avoid unnecessary font reloading
		drawTime(dc)

		dc.SavePNG("out.png")

		time.Sleep(time.Minute)
	}

}

func loadWeather(apiKey string) {
	// if lastWeatherCheck.IsZero() {
	// load weather
	// creates a new weather object with metric temperature and english as the returned language
	currentWeather, err := owm.NewCurrent("C", "EN", apiKey)
	if err != nil {
		panic(err)
	}
	// queries api and fills struct with data
	currentWeather.CurrentByCoordinates(coords)
	// fmt.Println(currentWeather)
	lastWeatherCheck = time.Now()
	// }

	// file, _ := ioutil.ReadFile("../weather.json")
	// _ = json.Unmarshal([]byte(file), &currentWeather)
}

func drawWeather(dc *gg.Context) {

	weatherWidth := float64(maxWidth) / 2
	// weatherHeight := float64(maxHeight) / 2

	desc := ""
	temp := fmt.Sprintf("%.2f°C", currentWeather.Main.Temp)
	minMax := fmt.Sprintf("%.2f° / %.2f°", currentWeather.Main.TempMin, currentWeather.Main.TempMax)
	iconPath := ""
	if len(currentWeather.Weather) > 0 {
		desc = currentWeather.Weather[0].Description
		_, err := owm.RetrieveIcon("../Icons", currentWeather.Weather[0].Icon+".png")
		if err != nil {
			panic(err)
		}
		iconPath = "../Icons/" + currentWeather.Weather[0].Icon + ".png"
	}
	if iconPath != "" {
		im, err := gg.LoadPNG(iconPath)
		if err != nil {
			panic(err)
		}
		image := imaging.Resize(im, 180, 180, imaging.CatmullRom)
		dc.DrawImage(image, 100, 210)
	}

	dc.DrawStringAnchored(desc, weatherWidth/2, 460, 0.5, 0.5)

	dc.DrawStringAnchored(minMax, weatherWidth/2, 500, 0.5, 0.5)

	fmt.Println(desc)
	fmt.Println(temp)
	fmt.Println(minMax)

	_ = dc.LoadFontFace(fontPath, 50)

	dc.DrawStringAnchored(temp, weatherWidth/2, 400, 0.5, 0.5)
}

func drawTime(dc *gg.Context) {
	currTime := time.Now()
	clock := currTime.Format("3:04 pm")
	_ = dc.LoadFontFace(fontPath, 90)

	dc.DrawStringAnchored(clock, maxWidth/2, maxHeight/5, 0.5, 0.5)
}

func drawDate(dc *gg.Context) {
	weekday := time.Now().Weekday()
	day := time.Now().Day()
	month := time.Now().Month()
	year := time.Now().Year()
	hour := time.Now().Hour()
	greeting := ""

	if hour >= 0 && hour < 12 {
		// 0 -> 11
		greeting = "Good Morning"
	} else if hour >= 12 && hour <= 17 {
		// 12 -> 5
		greeting = "Good Afternoon"
	} else if hour > 17 && hour < 20 {
		// 5-> 8
		greeting = "Good Evening"
	} else {
		// 8 -> midnight
		greeting = "Goodnight"
	}

	// example: Good Afternoon today is Mon 3 April 2023
	fullGreeting := fmt.Sprintf("%s today is %s %d %s %d", greeting, weekday, day, month, year)
	dc.DrawStringAnchored(fullGreeting, maxWidth/2, margin*2, 0.5, 0.5)
}

func drawHistory(dc *gg.Context, history HistoryToday) {
	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(history.Events))
	event := history.Events[index]
	formattedHistory := fmt.Sprintf("%s %s %s", history.Date, event.Year, event.Description)

	historyWidth := float64(maxWidth) / 2
	quarter := historyWidth / 2
	historyHeight := float64(maxHeight) / 2

	dc.DrawStringWrapped(formattedHistory, quarter*3, historyHeight, 0.5, 0.5, historyWidth-(margin*2), 1.5, gg.AlignLeft)
}
