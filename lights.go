package main

import (
	"log"
	"net/http"
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
	Bridge           *HueBridge
	ID               string
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
	username = "DQ3k2QsIs3lWJ8g4e3VmCzNBEIfi8vzfQmk25HwP"
	client   = &http.Client{}
)

func (l *Light) On() error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		On: true,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) Off() error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		On: false,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) IsOn() bool {
	return l.State.On
}

func (l *Light) Rename() error {
	return nil
}

func (l *Light) SetBri(v int) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Bri: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetHue(v int) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Hue: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetSat(v int) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Sat: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetXy(v []float32) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Xy: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetCt(v int) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Ct: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetTransitionDuration() error {
	return nil
}

func (l *Light) SetEffect(v string) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Effect: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetAlert(v string) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Alert: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetColorMode(v string) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		ColorMode: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (l *Light) SetMode(v string) error {
	_, err := l.Bridge.SetLightState(l.ID, State{
		Mode: v,
	})

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
