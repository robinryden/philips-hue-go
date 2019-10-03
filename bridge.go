package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func (b *HueBridge) GetLights() map[string]*Light {
	resp, err := http.Get(fmt.Sprint("http://", b.InternalIPAdress, "/api/", username, "/lights"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	lights := make(map[string]*Light)
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(r, &lights)
	if err != nil {
		log.Fatal(err)
	}

	for i, l := range lights {
		l.ID = string(i)
	}

	return lights
}

func (b *HueBridge) GetLightByID(id int) (*Light, error) {
	resp, err := http.Get(fmt.Sprint("http://", b.InternalIPAdress, "/api/", username, "/lights/", id))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	light := &Light{}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return light, err
	}

	err = json.Unmarshal(r, &light)
	if err != nil {
		log.Fatal(err)
		return light, err
	}

	return light, nil
}

func (b *HueBridge) SearchLights(light string) (bool, error) {
	resp, err := http.PostForm(fmt.Sprintf(b.InternalIPAdress, "/api/", username, "/lights"), url.Values{
		"deviceid": []string{b.ID},
	})
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	defer resp.Body.Close()

	return true, nil
}

func (b *HueBridge) SetLightState(id string, light State) (bool, error) {
	d := light
	json, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(json))

	req, err := http.NewRequest(http.MethodPut, fmt.Sprint("http://", b.InternalIPAdress, "/api/", username, "/lights/", id, "/state"), bytes.NewBuffer(json))
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}
