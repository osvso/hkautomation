package controller

import (
	"github.com/brutella/hc/accessory"
	"fmt"
	"hkautomation/rest"
)

type HKKitchenLightController struct {
	ServiceUrl string
	AccessoryStateUpdater rest.AccessoryStateUpdater
}

func (c HKKitchenLightController) changeLightState(newState int) {
	var accessoryName string = "kitchen_light"

	if newState == rest.LowState {
		fmt.Println("Turn kitchen light off")
	} else {
		fmt.Println("Turn kitchen light on")
	}

	action := rest.AccessoryStateAction{Name: accessoryName, State: newState}
	c.AccessoryStateUpdater.Update(action)
}

func (c HKKitchenLightController) Create() *accessory.Lightbulb {
	info := accessory.Info{
		Name:         "Kuchnia",
		Manufacturer: "osvso",
	}

	lightAcc := accessory.NewLightbulb(info)

	lightAcc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			c.changeLightState(rest.HighState)
		} else {
			c.changeLightState(rest.LowState)
		}
	})

	return lightAcc
}