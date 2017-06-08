package client

import (
	"fmt"
	"io"
	"net"

	"github.com/davidqhr/socccks/utils"
)

type Client struct {
	Id         string
	Conn       net.Conn
	AuthMethod byte
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewClient(conn net.Conn) *Client {
	return &Client{
		Conn: conn,
		Id:   randStringRunes(32),
	}
}

func (client *Client) GetSupportAuthMethods() (methods []byte, err error) {
	conn := client.Conn
	var buf = make([]byte, 100)
	_, er := io.ReadFull(conn, buf[:2])

	if er != nil {
		err = er
		return
	}

	version := buf[0]

	if version != utils.Version {
		err = fmt.Errorf("DO NOT SUPPORT PROXY Version %X", version)
		return
	}

	methodsCount := int(buf[1])

	_, er = io.ReadFull(conn, buf[:methodsCount])

	if er != nil {
		err = er
		return
	}

	methods = buf[:methodsCount]
	return
}

func (client *Client) SetAuthMethod(method byte) error {
	client.AuthMethod = method
	_, err := client.Conn.Write([]byte{utils.Version, method})
	return err
}

func (client *Client) AuthSuccess() error {
	_, err := client.Conn.Write([]byte("\x01\x00"))
	return err
}
