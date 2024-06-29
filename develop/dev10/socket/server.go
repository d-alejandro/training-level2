package socket

import (
	"fmt"
	"net"
	"sync"
	"unsafe"
)

type Server struct {
	listener  net.Listener
	quit      chan struct{}
	waitGroup sync.WaitGroup
}

func NewServer(network, address string) (*Server, error) {
	server := &Server{
		quit: make(chan struct{}),
	}

	var err error

	if server.listener, err = net.Listen(network, address); err != nil {
		return nil, err
	}

	fmt.Println("listening on", server.listener.Addr())

	server.waitGroup.Add(1)
	go server.serve()

	return server, nil
}

func (receiver *Server) Stop() error {
	close(receiver.quit)

	if err := receiver.listener.Close(); err != nil {
		return err
	}

	receiver.waitGroup.Wait()

	fmt.Println("\nstopped successfully")

	return nil
}

func (receiver *Server) serve() {
	defer receiver.waitGroup.Done()

	for {
		connection, err := receiver.listener.Accept()

		if err != nil {
			select {
			case <-receiver.quit:
				return
			default:
				fmt.Println(err)
				receiver.closeConnection(connection)
				continue
			}
		}

		receiver.waitGroup.Add(1)

		go func() {
			receiver.handleConnection(connection)
			receiver.waitGroup.Done()
		}()
	}
}

func (receiver *Server) handleConnection(connection net.Conn) {
	defer receiver.closeConnection(connection)

	var buffer []byte

	for {
		readLength, readError := connection.Read(buffer)
		if readError != nil {
			fmt.Println(readError)
			return
		}

		if readLength == 0 {
			return
		}

		inputString := string(buffer)

		fmt.Printf("received from %v: %s", connection.RemoteAddr(), inputString)

		outputString := fmt.Sprintf(
			"read length: %d, length: %d, size: %d",
			readLength,
			len(inputString),
			unsafe.Sizeof(inputString),
		)

		_, writeError := connection.Write([]byte(outputString))
		if writeError != nil {
			fmt.Println(writeError)
			return
		}
	}
}

func (receiver *Server) closeConnection(connection net.Conn) {
	if err := connection.Close(); err != nil {
		fmt.Println(err)
	}
}
