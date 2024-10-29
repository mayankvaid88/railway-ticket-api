package config

import (
	"encoding/json"
	"os"
)

type Config interface {
	GetSections() []Section
	GetTicketCost() float32
}

func NewConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		panic("unable to open the config file" + err.Error())
	}
	defer file.Close()
	var config config
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		panic("unable to decode the config file" + err.Error())
	}
	return config
}

type config struct {
	TicketCost float32	`json:"ticket_cost"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name          string `json:"name"`
	NumberOfSeats int  `json:"numberOfSeats"`
}

func (c config) GetSections() []Section {
	return c.Sections
}

func (c config) GetTicketCost() float32{
	return c.TicketCost
}
