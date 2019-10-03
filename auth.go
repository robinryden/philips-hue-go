package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Success struct {
		Username string `json:"username"`
	} `json:"success"`
}

func (bridge *HueBridge) CreateUser(device string) *Response {
	d := map[string]string{"devicetype": device}
	json, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(fmt.Sprint("http://", bridge.InternalIPAdress, "/api"), "application/json", bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dr := &Response{}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(r))

	return dr
}
