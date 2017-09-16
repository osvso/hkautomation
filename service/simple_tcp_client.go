package service

import (
	"net"
	"fmt"
	"bufio"
)

type SimpleTcpClient struct {
	AccessoryAuthority string
}

func (c SimpleTcpClient) Send(dataToSend string, opStateChannel chan<- bool) {
	conn, err := net.Dial("tcp", c.AccessoryAuthority)
	if err != nil {
		logConnectionError(err, c.AccessoryAuthority)
		opStateChannel <- false
		return
	}

	fmt.Fprintf(conn, dataToSend + "\n")

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		logConnectionError(err, c.AccessoryAuthority)
		opStateChannel <- false
	}

	if response == "ack" {
		fmt.Printf("Command acknowledged by %s", c.AccessoryAuthority)
		opStateChannel <- true
	} else {
		fmt.Printf("Command rejected by %s", c.AccessoryAuthority)
		opStateChannel <- false
	}

	conn.Close()
	return
}

func logConnectionError(err error, accessoryServerIp string) {
	if opErr, ok := err.(*net.OpError); ok {
		fmt.Printf("Failed to establish tcp connection with %s\n", accessoryServerIp)
		fmt.Printf("%#v\n", opErr)
	}
}
