package socket

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
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
	defer func() {
		if err := connection.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	remoteAddr := connection.RemoteAddr()

	reader := bufio.NewReader(connection)

	for {
		inputString, readError := reader.ReadString('\n')
		if readError != nil {
			if readError == io.EOF {
				fmt.Printf("client %v closed\n", remoteAddr)
				return
			}
			fmt.Println(readError)
			return
		}

		if inputString == "\n" {
			continue
		}

		inputString = strings.TrimSpace(inputString)

		fmt.Printf("received from %v: '%s'\n", remoteAddr, inputString)

		outputString := fmt.Sprintf("read length: %d, size: %d\n", len(inputString), unsafe.Sizeof(inputString))

		_, writeError := connection.Write([]byte(outputString))
		if writeError != nil {
			fmt.Println(writeError)
			return
		}
	}
}
