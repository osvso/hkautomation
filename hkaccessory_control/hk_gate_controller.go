package hkaccessory_control

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"hkautomation/service"
	"hkautomation/hkaccessory"
	"hkautomation/model"
)

type HKGateController struct {
	ServiceUrl            string
	AccessoryStateUpdater service.AccessoryStateUpdater
}

func (c HKGateController) Create() *hkaccessory.GarageDoor {
	info := accessory.Info{
		Name:         "Brama wjazdowa",
		Manufacturer: "osvso",
		Model:        "01",
	}

	gateAcc := hkaccessory.NewGarageDoor(info)

	gateAcc.GarageDoor.TargetDoorState.OnValueRemoteUpdate(func(newValue int) {
		var opStateChannel chan int = make(chan int)
		go c.updateState(gateAcc, opStateChannel)

		if newValue == characteristic.TargetDoorStateOpen {
			go c.AccessoryStateUpdater.Update(model.LowState, &info, opStateChannel)
		} else if newValue == characteristic.TargetDoorStateClosed {
			go c.AccessoryStateUpdater.Update(model.HighState, &info, opStateChannel)
		}
	})

	return gateAcc
}

func (c HKGateController) updateState(gateAcc *hkaccessory.GarageDoor, opStateChannel <-chan int) {
	newState := <-opStateChannel
	gateAcc.GarageDoor.CurrentDoorState.SetValue(newState)
}