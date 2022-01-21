package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ClockResponse struct {
	Status string    `json:"status"`
	Data   ClockData `json:"data"`
}

type ClockData struct {
	Modules map[string]ClockModule `json:"modules"`
}

type ClockModule struct {
	Type           string    `json:"type"`
	Flavor         string    `json:"flavor"`
	Description    string    `json:"description"`
	UpdateInterval int       `json:"update_interval_seconds"`
	Labels         []string  `json:"labels,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
}

func query() (ClockResponse, error) {
	url := "https://api.climateclock.world/v1/clock"
	var response ClockResponse

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return response, err
	}

	req.Header.Set("User-Agent", "climate-clock-tui")

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&response); err != nil {
		return response, err
	}

	return response, nil
}
