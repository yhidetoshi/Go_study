package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	url   = "https://api.nature.global/1/devices"
	token = ""
)

type Device struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	TemperatureOffset int32        `json:"temperature_offset"`
	HumidityOffset    int32        `json:"humidity_offset"`
	CreatedAt         string       `json:"created_at"`
	UpdatedAt         string       `json:"updated_at"`
	FirmwareVersion   string       `json:"firmware_version"`
	NewestEvents      NewestEvents `json:"newest_events"`
}

type NewestEvents struct {
	Temperature SensorValue `json:"te"`
	Humidity    SensorValue `json:"hu"`
	Illuminance SensorValue `json:"il"`
}

type SensorValue struct {
	Value     float64 `json:"val"`
	CreatedAt string  `json:"created_at"`
}

func (d *Device) FetchValesFromNatureRemo() {
	var devices []*Device

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error")
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error request")
	}

	err = json.NewDecoder(resp.Body).Decode(&devices)
	if err != nil {
		fmt.Println("Error decode")
	}
	fmt.Printf("%v\t%v\t%v\n",
		devices[0].NewestEvents.Temperature.Value,
		devices[0].NewestEvents.Humidity.Value,
		devices[0].NewestEvents.Illuminance.Value)
}

func main() {
	d := &Device{}
	d.FetchValesFromNatureRemo()
}
