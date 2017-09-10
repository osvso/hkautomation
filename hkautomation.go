package main

import (
	"github.com/brutella/hc"
	"fmt"
	"github.com/brutella/hc/accessory"
	"homeAutomation/controller"
	"homeAutomation/rest"
)

const CONTROLLER_IP_ADDRESS string = "192.168.0.166"
const PROTOCOL string = "http://"
const GATE_SERVICE_URL = PROTOCOL + CONTROLLER_IP_ADDRESS

var restClient rest.SimpleRestClient = rest.SimpleRestClient{}

func createGateAccessory() *accessory.Accessory {
	gateAccessory := controller.HKGateController{ServiceUrl: GATE_SERVICE_URL, RestClient: restClient}.Create()

	return gateAccessory.Accessory
}

func createKitchenLightAccessory() *accessory.Accessory {
	kitchenLightAccessory := controller.HKKitchenLightController{ServiceUrl: GATE_SERVICE_URL, RestClient: restClient}.Create()

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
