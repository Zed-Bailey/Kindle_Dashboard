package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HistoryToday struct {
	Date      string         `json:"date"`
	Wikipedia string         `json:"wikipedia"`
	Events    []HistoryEvent `json:"events"`
}

type HistoryEvent struct {
	Year        string `json:"year"`
	Description string `json:"description"`

	// Note: optional
	// Wikipedia   []struct {
	// 	Title     string `json:"title"`
	// 	Wikipedia string `json:"wikipedia"`
	// } `json:"wikipedia"`
}

func GetHistory() (HistoryToday, error) {

	url := fmt.Sprintf("https://byabbe.se/on-this-day/%d/%d/events.json", time.Now().Month(), time.Now().Day())
	resp, err := http.Get(url)
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
