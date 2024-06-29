package socket

import (
	"bufio"
	"net"
	"strings"
	"time"
)

type Client struct {
	connection net.Conn
	readWriter *bufio.ReadWriter
}

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

func (receiver *Client) Send(message string) (string, error) {
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

func (receiver *Client) Stop() error {
	if err := receiver.connection.Close(); err != nil {
		return err
	}
	return nil
}
