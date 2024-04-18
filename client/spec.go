package client

import (
	"errors"
	"log/slog"
	"net"
)

type Client struct {
	conn *net.UDPConn
}

func New() *Client {
	return &Client{
		conn: nil,
	}, nil
}

func (c *Client) Connect(address string) (final error) {
	server, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return err
	}

	c.conn, err = net.DialUDP("udp4", nil, server)
	if err != nil {
		return err
	}

	defer func() {
		// If discovery fails, the connection is closed
		if final != nil {
			final = errors.Join(final, c.conn.Close())
		}
	}()

	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		slog.Debug("disconnecting")
		return c.conn.Close()
	}
	return nil
}
