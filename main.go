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
		//fmt.Println(bridge.CreateUser("my-hue-bridge"))

		lights := bridge.GetLights()
		fmt.Println(lights)

		for _, l := range lights {
			//l.On()
			l.On()
			fmt.Println(l.IsOn())
		}
	}
}
