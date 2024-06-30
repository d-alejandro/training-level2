package socket

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type ConnectionController struct {
	connection  net.Conn
	quitChannel chan struct{}
	once        sync.Once
}

func NewConnectionController(host, port string, timeout time.Duration) (*ConnectionController, error) {
	const NetworkProtocolTCP = "tcp"

	connectionController := &ConnectionController{
		quitChannel: make(chan struct{}),
	}

	var err error

	address := net.JoinHostPort(host, port)
	connectionController.connection, err = net.DialTimeout(NetworkProtocolTCP, address, timeout)

	if err != nil {
		return nil, err
	}

	return connectionController, nil
}

func (receiver *ConnectionController) CheckConnection() <-chan struct{} {
	go func() {
		buffer := make([]byte, 10)

		if _, err := receiver.connection.Read(buffer); err != nil {
			select {
			case <-receiver.quitChannel:
				fmt.Println("use of closed network connection.")
				return
			default:
				if err == io.EOF {
					fmt.Println("server closed.")
					receiver.Stop()
					return
				}
				fmt.Println("check connection controller:", err)
			}
		}
	}()

	return receiver.quitChannel
}

func (receiver *ConnectionController) Stop() {
	receiver.once.Do(func() {
		close(receiver.quitChannel)

		if err := receiver.connection.Close(); err != nil {
			fmt.Println("close connection controller:", err)
		}
	})
}
