package main

import (
	"fmt"
	"log"
)

func main() {
	bridges, err := GetBridges()
	if err != nil {
		log.Fatal(err)
	}

	for _, bridge := range bridges {
		fmt.Println(bridge.InternalIPAdress)
		//bridge.CreateUser("my-hue-bridge")

		lights := bridge.GetLights()

		for _, l := range lights {
			l.SetLightState(bridge, "3", &State{
				On: false,
			})
			fmt.Println(l.State.On)
		}
	}
}
