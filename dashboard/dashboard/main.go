package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// kindle 4 horizontal screen size
const width int = 800
const height int = 600

type Config struct {
	Name              string       `json:"name"`
	Lat               float64      `json:"lat"`
	Lon               float64      `json:"lon"`
	OpenweatherApiKey string       `json:"openweatherApiKey"`
	LastUpdate        time.Time    `json:"lastUpdate"`
	HistoryToday      HistoryToday `json:"historyToday"`
}

func displayError(e string) {
	// fmt.Println("Error: " + e) // debug
	// exec.Command("eips", fmt.Sprintf("%d", width/2), fmt.Sprintf("%d", height/2), e)
	log.Println(e)
}

func main() {

	// load config file
	configFile, err := ioutil.ReadFile("../config.json")
	if err != nil {
		displayError("couldn't read config.json file")
		return
	}

	config := Config{}
	_ = json.Unmarshal(configFile, &config)

	// load template svg
	data, err := os.ReadFile("../svg/template.svg")

	if err != nil {
		displayError("couldn't read template.svg")
		return
	}

	dataString := string(data)
	updatedTemplate := updateInterval(&config, dataString)

	// write the updated template to a new file
	// os.WriteFile("./temp.svg", []byte(updatedTemplate), os.ModeAppend) // debug
	os.WriteFile("../svg/temp.svg", []byte(updatedTemplate), os.ModeAppend)
	json, _ := json.Marshal(&config)
	os.WriteFile("../config.json", json, os.ModeAppend)
}

// updates the time in the svg
func updateInterval(config *Config, svgData string) string {
	currTime := time.Now()
	clock := currTime.Format("3:04 pm")

	svgData = strings.Replace(svgData, "%TIME%", clock, 1)

	if time.Since(config.LastUpdate) >= time.Minute*10 {
		fmt.Println("more then 10 minutes since last update") // debug
		// update weather
		// update history today
		svgData = strings.Replace(svgData, "%TMP%", "24°", 1)
		svgData = strings.Replace(svgData, "%condition%", "cloudy", 1)
	}

	if currTime.Hour() == 0 || config.HistoryToday.Events == nil {
		// fetch new history today
		history, err := GetHistory()
		if err != nil {
			svgData = strings.Replace(svgData, "%history_today%", err.Error(), 1)
		} else {
			config.HistoryToday = history
		}

	}

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(config.HistoryToday.Events))
	event := config.HistoryToday.Events[index]
	formattedHistory := fmt.Sprintf("%s %s %s", config.HistoryToday.Date, event.Year, event.Description)
	wordwrap := WordWrap(formattedHistory, 70)
	fmt.Println(wordwrap)
	svgData = strings.Replace(svgData, "%history_today%", wordwrap, 1)

	return svgData
}
