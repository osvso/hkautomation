package main

import (
	"github.com/brutella/hc"
	"fmt"
	"github.com/brutella/hc/accessory"
	"hkautomation/service"
	"hkautomation/hkaccessory_control"
)

const ARDUINO_AUTHORITY string = "192.168.0.166:5050"

var accessoryStateUpdater *service.AccessoryStateUpdater = service.NewAccessoryStateUpdater(ARDUINO_AUTHORITY)

func createGateAccessory() *accessory.Accessory {
	gateAccessory := hkaccessory_control.HKGateController{ServiceUrl: ARDUINO_AUTHORITY, AccessoryStateUpdater: *accessoryStateUpdater}.Create()
	return gateAccessory.Accessory
}

func createKitchenLightAccessory() *accessory.Accessory {
	kitchenLightAccessory := hkaccessory_control.HKKitchenLightController{ServiceUrl: ARDUINO_AUTHORITY, AccessoryStateUpdater: *accessoryStateUpdater}.Create()

	return kitchenLightAccessory.Accessory
}

func main() {
	t, err := hc.NewIPTransport(hc.Config{Pin: "00102003"}, createGateAccessory(), createKitchenLightAccessory())
	if err != nil {
		fmt.Println(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
