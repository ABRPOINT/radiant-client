package client

import (
	"encoding/json"
	"net"
	"radiant/protocol"
)

type Client struct {
	conn net.Conn
}

func NewClient(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: conn,
	}

	return client, nil
}

func (c *Client) SendMessage(message string) error {
	msg := protocol.Message{
		Type:    protocol.MessageTypeSendMessage,
		Payload: message,
	}

	return c.send(msg)
}

func (c *Client) send(message protocol.Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(data)
	return err
}
