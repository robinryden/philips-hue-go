package main

import (
	"log"
	"testing"
)

func TestOn(t *testing.T) {
	bridges, err := GetBridges()
	if err != nil {
		t.Fatal(err)
	}

	l, err := bridges[0].GetLightByID(1)
	if err != nil {
		t.Fatal(err)
	}

	err = l.On()
	if err != nil {
		log.Fatal(err)
	}

	t.Logf("Turned on light %s", l.ID)
}
