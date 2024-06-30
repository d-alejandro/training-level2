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
	waitGroup   sync.WaitGroup
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

	connectionController.waitGroup.Add(1)
	go connectionController.checkConnection()

	return connectionController, nil
}

func (receiver *ConnectionController) GetQuitChannel() <-chan struct{} {
	return receiver.quitChannel
}

func (receiver *ConnectionController) Stop() {
	receiver.closeQuitChannel()

	if err := receiver.connection.Close(); err != nil {
		select {
		case <-receiver.GetQuitChannel():
			break
		default:
			fmt.Println("close connection controller:", err)
		}
	}

	receiver.waitGroup.Wait()
}

func (receiver *ConnectionController) checkConnection() {
	defer receiver.waitGroup.Done()

	buffer := make([]byte, 10)

	if _, err := receiver.connection.Read(buffer); err != nil {
		select {
		case <-receiver.quitChannel:
			fmt.Println("closing the network connection.")
			return
		default:
			if err == io.EOF {
				fmt.Println("server closed.")
				receiver.closeQuitChannel()
				return
			}
			fmt.Println("check connection controller:", err)
		}
	}
}

func (receiver *ConnectionController) closeQuitChannel() {
	receiver.once.Do(func() {
		close(receiver.quitChannel)
	})
}
