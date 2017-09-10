package main

import (
	"github.com/brutella/hc"
	"fmt"
	"github.com/brutella/hc/accessory"
	"hkautomation/rest"
	"hkautomation/controller"
)

const CONTROLLER_IP_ADDRESS string = "192.168.0.166"
const PROTOCOL string = "http://"
const ARDUINO_SERVICE_URL = PROTOCOL + CONTROLLER_IP_ADDRESS

var accessoryStateUpdater rest.AccessoryStateUpdater = rest.AccessoryStateUpdater{ServiceUrl: ARDUINO_SERVICE_URL}

func createGateAccessory() *accessory.Accessory {
	gateAccessory := controller.HKGateController{ServiceUrl: ARDUINO_SERVICE_URL, AccessoryStateUpdater: accessoryStateUpdater}.Create()

	return gateAccessory.Accessory
}

func createKitchenLightAccessory() *accessory.Accessory {
	kitchenLightAccessory := controller.HKKitchenLightController{ServiceUrl: ARDUINO_SERVICE_URL, AccessoryStateUpdater: accessoryStateUpdater}.Create()

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
