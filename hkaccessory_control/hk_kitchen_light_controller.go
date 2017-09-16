package hkaccessory_control

import (
	"github.com/brutella/hc/accessory"
	"hkautomation/service"
	"hkautomation/model"
)

type HKKitchenLightController struct {
	ServiceUrl            string
	AccessoryStateUpdater service.AccessoryStateUpdater
}

func (c HKKitchenLightController) Create() *accessory.Lightbulb {
	info := accessory.Info{
		Name:         "Kuchnia",
		Manufacturer: "osvso",
		Model:        "02",
	}

	lightAcc := accessory.NewLightbulb(info)

	lightAcc.Lightbulb.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			go c.AccessoryStateUpdater.Update(model.HighState, &info, make(chan bool))
		} else {
			go c.AccessoryStateUpdater.Update(model.LowState, &info, make(chan bool))
		}
	})

	return lightAcc
}
