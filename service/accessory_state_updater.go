package service

import (
	"fmt"
	"github.com/brutella/hc/accessory"
	"strconv"
)


type AccessoryStateUpdater struct {
	simpleTcpClient SimpleTcpClient
}

func NewAccessoryStateUpdater(AccessoryAuthority string) *AccessoryStateUpdater {
	client := SimpleTcpClient{AccessoryAuthority:AccessoryAuthority}
	return &AccessoryStateUpdater{client}
}

func (u AccessoryStateUpdater) Update(newState int, info *accessory.Info, opStateChannel chan<- int) {
	fmt.Printf("Changing accessory '%s' state to %d\n", info.Name, newState)

	command := u.createCommand(newState, info)
	u.simpleTcpClient.Send(command, opStateChannel)
}

func (u AccessoryStateUpdater) createCommand(newState int, info *accessory.Info) string {
	var command string = info.Model + ":" + strconv.Itoa(newState)
	fmt.Printf("Sending command %s\n", command)
	return command
}
