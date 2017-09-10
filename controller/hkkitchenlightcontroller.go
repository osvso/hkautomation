package controller

import (
	"github.com/brutella/hc/accessory"
	"fmt"
	"homeAutomation/rest"
)

type HKKitchenLightController struct {
	ServiceUrl string
	RestClient rest.SimpleRestClient
}

func (c HKKitchenLightController) turnLightOn() {
	fmt.Println("Turn kitchen light on")
	c.RestClient.Post(c.ServiceUrl, "{\"kitchen_light\":\"1\"}")
}

func (c HKKitchenLightController) turnLightOff() {
	fmt.Println("Turn kitchen light off")
	c.RestClient.Post(c.ServiceUrl, "{\"kitchen_light\":\"0\"}")
}

func (c HKKitchenLightController) Create() *accessory.Lightbulb {
	info := accessory.Info{
		Name:         "Kuchnia",
		Manufacturer: "osvso",
	}

	lightAcc := accessory.NewLightbulb(info)

	lightAcc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			c.turnLightOn()
		} else {
			c.turnLightOff()
		}
	})

	return lightAcc
}