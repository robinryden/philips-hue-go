package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Light struct {
	State            *State        `json:"state,omitempty"`
	SwUpdate         *SwUpdate     `json:"swupdate,omitempty"`
	Type             string        `json:"type,omitempty"`
	Name             string        `json:"name,omitempty"`
	ModelID          string        `json:"modelid,omitempty"`
	ManufacturerName string        `json:"manufacturername,omitempty"`
	ProductName      string        `json:"productname,omitempty"`
	Capabilities     *Capabilities `json:"capabilities,omitempty"`
	Config           *Config       `json:"config,omitempty"`
	UniqueID         string        `json:"uniqueid,omitempty"`
	SwVersion        string        `json:"swversion,omitempty"`
}

type State struct {
	On        bool      `json:"on"`
	Bri       int       `json:"bri,omitempty"`
	Hue       int       `json:"hue,omitempty"`
	Sat       int       `json:"sat,omitempty"`
	Effect    string    `json:"effect,omitempty"`
	Xy        []float32 `json:"xy,omitempty"`
	Ct        int       `json:"ct,omitempty"`
	Alert     string    `json:"alert,omitempty"`
	ColorMode string    `json:"colormode,omitempty"`
	Mode      string    `json:"mode,omitempty"`
	Reachable bool      `json:"reachable,omitempty"`
}

type SwUpdate struct {
	State       string `json:"state,omitempty"`
	LastInstall string `json:"lastinstall,omitempty"`
}

type Capabilities struct {
	Certified bool     `json:"certified,omitempty"`
	Control   *Control `json:"control,omitempty"`
	Config    *Config  `json:"config,omitempty"`
	UniqueID  string   `json:"uniqueid,omitempty"`
	SwVersion string   `json:"swversion,omitempty"`
}

type Control struct {
	MinDimLevel    int         `json:"mindimlevel,omitempty"`
	MaxLumen       int         `json:"maxlumen,omitempty"`
	ColorGamutType string      `json:"colorgamuttype,omitempty"`
	ColorGamut     [][]float32 `json:"colorgamut,omitempty"`
	Ct             *Ct         `json:"ct,omitempty"`
	Streaming      *Streaming  `json:"streaming,omitempty"`
}

type Ct struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

type Streaming struct {
	Renderer bool `json:"renderer,omitempty"`
	Proxy    bool `json:"proxy,omitempty"`
}

type Config struct {
	Archetype string `json:"archetype,omitempty"`
	Function  string `json:"function,omitempty"`
	Direction string `json:"direction,omitempty"`
}

var (
	username = "tOLctHltX1lJQKmG15wI71QOohiPDzGHb3BhCxd6"
)

func (bridge *HueBridge) GetLights() map[string]Light {
	resp, err := http.Get(fmt.Sprint("http://", bridge.InternalIPAdress, "/api/", username, "/lights"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	lights := map[string]Light{}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(r, &lights)
	if err != nil {
		log.Fatal(err)
	}

	return lights
}

func (bridge *HueBridge) GetLightByID(id int) *Light {
	resp, err := http.Get(fmt.Sprintf(bridge.InternalIPAdress, "/api/", username, "/lights/", id))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	light := &Light{}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(r, &light)
	if err != nil {
		log.Fatal(err)
	}

	return light
}

func (bridge *HueBridge) SearchLights(light string) (bool, error) {
	resp, err := http.PostForm(fmt.Sprintf(bridge.InternalIPAdress, "/api/", username, "/lights"), url.Values{
		"deviceid": []string{bridge.ID},
	})
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	defer resp.Body.Close()

	return true, nil
}
