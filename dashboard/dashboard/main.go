package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// kindle 4 horizontal screen size
const width int = 800
const height int = 600

type Config struct {
	Name              string    `json:"name"`
	Lat               float64   `json:"lat"`
	Lon               float64   `json:"lon"`
	OpenweatherApiKey string    `json:"openweatherApiKey"`
	LastUpdate        time.Time `json:"lastUpdate"`
	HistoryToday      []string  `json:"historyToday"`
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
}

// updates the time in the svg
func updateInterval(config *Config, svgData string) string {
	currTime := time.Now()
	clock := currTime.Format("3:04 pm")

	svgData = strings.Replace(svgData, "%TIME%", clock, 1)
	return svgData

	// if time.Since(config.LastUpdate) >= time.Minute*10 {
	// 	fmt.Println("more then 10 minutes since last update") // debug
	// 	// update weather
	// 	// update history today
	// }

	// if time.Now().Hour() == 0 {
	// 	// fetch new history today
	// }
}
