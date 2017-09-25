package service

import (
	"net"
	"fmt"
	"bufio"
)

type SimpleTcpClient struct {
	AccessoryAuthority string
}

func (c SimpleTcpClient) Send(dataToSend string, opStateChannel chan<- int) {
	conn, err := net.Dial("tcp", c.AccessoryAuthority)
	if err != nil {
		logConnectionError(err, c.AccessoryAuthority)
		opStateChannel <- 0
		return
	}

	fmt.Fprintf(conn, dataToSend + "\n")

	response, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		logConnectionError(err, c.AccessoryAuthority)
		opStateChannel <- 0
	}

	if len(response) == 4 {
		if response[2] ==  48 {
			fmt.Println("Command returned 0")
			opStateChannel <- 0
		} else {
			fmt.Println("Command returned 1")
			opStateChannel <- 1
		}
	} else {
		fmt.Println("Command returned 0")
		opStateChannel <- 0
	}

	conn.Close()
}

func logConnectionError(err error, accessoryServerIp string) {
	if opErr, ok := err.(*net.OpError); ok {
		fmt.Printf("Failed to establish tcp connection with %s\n", accessoryServerIp)
		fmt.Printf("%#v\n", opErr)
	}
}
