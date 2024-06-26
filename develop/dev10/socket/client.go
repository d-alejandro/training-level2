package socket

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

/*
Client structure
*/
type Client struct {
	connection net.Conn
	readWriter *bufio.ReadWriter
}

/*
NewClient constructor
*/
func NewClient(host, port string, timeout time.Duration) (*Client, error) {
	const NetworkProtocolTCP = "tcp"

	client := new(Client)

	var err error

	address := net.JoinHostPort(host, port)
	client.connection, err = net.DialTimeout(NetworkProtocolTCP, address, timeout)

	if err != nil {
		return nil, err
	}

	client.readWriter = bufio.NewReadWriter(
		bufio.NewReader(client.connection),
		bufio.NewWriter(client.connection),
	)

	return client, nil
}

/*
Send method
*/
func (receiver *Client) Send(message string) (string, error) {
	if message == "" {
		return "", errors.New("message is empty")
	}

	if _, err := receiver.readWriter.WriteString(message + "\n"); err != nil {
		return "", err
	}

	if err := receiver.readWriter.Flush(); err != nil {
		return "", err
	}

	readString, err := receiver.readWriter.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(readString), nil
}

/*
Stop method
*/
func (receiver *Client) Stop() {
	if err := receiver.connection.Close(); err != nil {
		fmt.Println(err)
	}
}
