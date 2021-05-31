package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Connection struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"connection"`
}

func initConfiguration() (*Configuration, error) {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	var config Configuration

	json.Unmarshal(file, &config)

	if config.Connection.Host == "" {
		config.Connection.Host = "0.0.0.0"
	}

	return &config, nil
}
