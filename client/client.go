package client

import (
	"net"
)

type Client struct {
	conn net.Conn
}

func (this *Client) quit() {
	this.conn.Close()
}
