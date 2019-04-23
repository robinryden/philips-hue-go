package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	discoverURL = "https://discovery.meethue.com"
)

type HueBridge struct {
	ID               string `json:"id"`
	InternalIPAdress string `json:"internalipaddress"`
	Username         string
}

func GetBridges() ([]HueBridge, error) {
	resp, err := http.Get(discoverURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	hb := []HueBridge{}
	bridges, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bridges, &hb)
	if err != nil {
		log.Fatal(err)
	}

	return hb, nil
}
