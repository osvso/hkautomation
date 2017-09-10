package rest

import (
	"homeAutomation/rest"
	"encoding/json"
)

var simpleRestClient rest.SimpleRestClient = rest.SimpleRestClient{}

type AccessoryStateUpdater struct {
	ServiceUrl string
}

func (u AccessoryStateUpdater) Update(action AccessoryStateAction) {
	marshaled, _ := json.Marshal(action)
	simpleRestClient.Post(u.ServiceUrl, string(marshaled))
}
