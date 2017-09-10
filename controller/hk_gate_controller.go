package controller

import (
	"github.com/brutella/hc/accessory"
	"fmt"
	"homeAutomation/hkaccessory"
	"github.com/brutella/hc/characteristic"
	"hkautomation/rest"
)

type HKGateController struct {
	ServiceUrl string
	AccessoryStateUpdater rest.AccessoryStateUpdater
}

func (c HKGateController) changeGateState(newState int) {
	var accessoryName string = "gate"

	if newState == rest.LowState {
		fmt.Println("Closing gate")
	} else {
		fmt.Println("Opening gate")
	}

	action := rest.AccessoryStateAction{Name: accessoryName, State: newState}
	c.AccessoryStateUpdater.Update(action)
}

func (c HKGateController) Create() *hkaccessory.GarageDoor {
	gateInfo := accessory.Info{
		Name:         "Brama wjazdowa",
		Manufacturer: "osvso",
	}

	gateAcc := hkaccessory.NewGarageDoor(gateInfo)

	gateAcc.GarageDoor.TargetDoorState.OnValueRemoteUpdate(func(newValue int) {
		if newValue == characteristic.TargetDoorStateOpen {
			c.changeGateState(rest.HighState)
		} else if newValue == characteristic.TargetDoorStateClosed {
			c.changeGateState(rest.LowState)
		}

		// for now the gate changes it's state immediately, can be changed afterwards and listen for an Arduino request
		gateAcc.GarageDoor.CurrentDoorState.SetValue(newValue)
	})

	return gateAcc
}
