package rest

import (
	"bytes"
	"io/ioutil"
	"fmt"
	"net/http"
)

type SimpleRestClient struct {
}

func (client SimpleRestClient) Post(requestUrl, jsonBody string) {
	body := bytes.NewBufferString(jsonBody)
	response, err := http.Post(requestUrl, "application/json", body)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(responseBody))
}
