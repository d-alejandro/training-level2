package socket

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
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

		fmt.Printf("client %v connected\n", connection.RemoteAddr())

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

	readWriter := bufio.NewReadWriter(
		bufio.NewReader(connection),
		bufio.NewWriter(connection),
	)

	const DeadlineTimeSecond = 15 * time.Second

	for {
		select {
		case <-receiver.quit:
			return
		default:
		}

		deadlineTime := time.Now().Add(DeadlineTimeSecond)
		if err := connection.SetDeadline(deadlineTime); err != nil {
			return
		}

		inputString, readError := readWriter.ReadString('\n')
		if readError != nil {
			var opError *net.OpError

			if readError == io.EOF {
				fmt.Printf("client %v closed\n", remoteAddr)
				return
			} else if errors.As(readError, &opError) && opError.Timeout() {
				fmt.Println(readError)
				continue
			}
			fmt.Println(readError)
			return
		}

		if inputString == "\n" {
			return
		}

		inputString = strings.TrimSpace(inputString)

		fmt.Printf("received from %v: '%s'\n", remoteAddr, inputString)

		outputString := fmt.Sprintf("read length: %d, size: %d\n", len(inputString), unsafe.Sizeof(inputString))

		if _, err := readWriter.WriteString(outputString); err != nil {
			fmt.Println(err)
			return
		}

		if err := readWriter.Flush(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
