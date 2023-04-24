package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type Conn struct {
	*nats.Conn
}

func NewConn(url string) (*Conn, error) {
	opts := nats.Options{
		Servers: []string{url},
	}

	nc, err := opts.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS server: %v", err)
	}

	return &Conn{nc}, nil
}

func (nc *Conn) Close() error {
	nc.Conn.Close()
	return nil
}

