package controller

import (
	"github.com/brutella/hc/accessory"
	"fmt"
	"homeAutomation/hkaccessory"
	"github.com/brutella/hc/characteristic"
	"homeAutomation/rest"
)

type HKGateController struct {
	ServiceUrl string
	RestClient rest.SimpleRestClient
}

func (c HKGateController) openGate() {
	fmt.Println("Opening gate")
	c.RestClient.Post(c.ServiceUrl, "{\"gate\":\"1\"}")
}

func (c HKGateController) closeGate() {
	fmt.Println("Closing gate")
	c.RestClient.Post(c.ServiceUrl, "{\"gate\":\"0\"}")
}

func (c HKGateController) Create() *hkaccessory.GarageDoor {
	gateInfo := accessory.Info{
		Name:         "Brama wjazdowa",
		Manufacturer: "osvso",
	}

	gateAcc := hkaccessory.NewGarageDoor(gateInfo)

	gateAcc.GarageDoor.TargetDoorState.OnValueRemoteUpdate(func(newValue int) {
		if newValue == characteristic.TargetDoorStateOpen {
			c.openGate()
		} else if newValue == characteristic.TargetDoorStateClosed {
			c.closeGate()
		}

		// for now the gate changes it's state immediately, can be changed afterwards and listen for an Arduino request
		gateAcc.GarageDoor.CurrentDoorState.SetValue(newValue)
	})

	return gateAcc
}
