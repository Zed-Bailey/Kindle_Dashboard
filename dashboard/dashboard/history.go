package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HistoryToday struct {
	Date      string `json:"date"`
	Wikipedia string `json:"wikipedia"`
	Events    []struct {
		Year        int    `json:"year"`
		Description string `json:"description"`
		Wikipedia   []struct {
			Title     string `json:"title"`
			Wikipedia string `json:"wikipedia"`
		} `json:"wikipedia"`
	} `json:"events"`
}

func GetHistory() (HistoryToday, error) {

	resp, err := http.Get("https://byabbe.se/on-this-day/1/12/events.json")
	if err != nil {
		return HistoryToday{}, err
	}

	// read body
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// parse json
	history := HistoryToday{}
	err = json.Unmarshal(body, &history)
	if err != nil {
		return HistoryToday{}, err
	}

	return history, nil
}
