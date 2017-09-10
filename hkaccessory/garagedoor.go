package hkaccessory

import (
	"github.com/brutella/hc/service"
	"github.com/brutella/hc/accessory"
)

type GarageDoor struct {
	*accessory.Accessory
	GarageDoor *service.GarageDoorOpener
}

func NewGarageDoor(info accessory.Info) *GarageDoor {
	acc := GarageDoor{}
	acc.Accessory = accessory.New(info, accessory.TypeGarageDoorOpener)
	acc.GarageDoor = service.NewGarageDoorOpener()

	acc.AddService(acc.GarageDoor.Service)

	return &acc
}
