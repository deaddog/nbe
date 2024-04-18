package client

import (
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"

	"github.com/deaddog/nbe/frames"
)

type Client struct {
	logger *slog.Logger
	conn   *net.UDPConn
	appId  frames.AppId
	serial frames.Serial
	next   frames.MessageId
}

func New(appId frames.AppId, logger *slog.Logger) (*Client, error) {
	if err := appId.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		logger: logger.With("nbe-app", appId),
		conn:   nil,
		appId:  appId,
		serial: 0,
		next:   0,
	}, nil
}

func (c *Client) Connect(address string) (final error) {
	c.logger.Debug("connecting")

	server, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return err
	}

	c.conn, err = net.DialUDP("udp4", nil, server)
	if err != nil {
		return err
	}

	c.logger.Debug("connected")

	defer func() {
		// If discovery fails, the connection is closed
		if final != nil {
			final = errors.Join(final, c.conn.Close())
		}
	}()

	payload, err := c.SendSync(frames.FunctionDiscovery, "NBE Discovery")
	if err != nil {
		return err
	}

	if serialStr, ok := payload["Serial"]; ok {
		serial, err := strconv.Atoi(serialStr)
		if err != nil {
			return fmt.Errorf("failed parsing serial '%s' in discovery payload: %w", serialStr, err)
		}
		c.serial = frames.Serial(serial)
	} else {
		return fmt.Errorf("missing Serial in discovery payload")
	}

	c.logger.Debug("identified controller", "serial", c.serial)

	return nil
}

func (c *Client) Close() error {
	if c.conn != nil {
		c.logger.Debug("disconnecting")
		return c.conn.Close()
	}
	return nil
}

// SendSync sends a message and waits for a response
func (c *Client) SendSync(function frames.Function, payload string) (frames.ResponsePayload, error) {
	c.logger.Debug("sending message", "id", c.next, "func", function, "payload", payload)

	req := frames.Request{
		AppId:          c.appId,
		Serial:         c.serial,
		EncryptionMode: frames.EncryptionModeNone,
		Function:       function,
		Id:             c.next,
		Password:       "",
		Timestamp:      time.Now().Truncate(time.Second),
		Payload:        payload,
	}

	c.next = (c.next + 1) % 100

	if err := req.Validate(); err != nil {
		return nil, err
	}

	data := frames.EncodeRequest(req)

	_, err := c.conn.Write(data)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("request sent")

	buffer := make([]byte, 4096)
	resp, err := frames.DecodeResponse(buffer)
	if err != nil {
		return nil, err
	}

	if resp.Code != frames.ResponseCodeOk {
		return nil, fmt.Errorf("response code was %v, expected %v (OK)", resp.Code, frames.ResponseCodeOk)
	}

	c.logger.Debug("response received and decoded")

	return resp.Payload, nil
}
