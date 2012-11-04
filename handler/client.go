//client

package handler

import (
	"crypto/cipher"
	"errors"
	"log"
	"net"
)

type Client struct {
	Server                     string
	ClientCipher, ServerCipher cipher.Block
}

func handshake(conn net.Conn) (err error) {
	b := make([]byte, 2)
	_, err = conn.Read(b)
	if err != nil {
		return
	}

	if b[0] != 5 {
		return errors.New("socks version incorrect")
	}

	b = make([]byte, b[1]) //auth methods
	conn.Read(b)
	conn.Write([]byte("\x05\x00")) //no auth
	return
}

func (c *Client) Handle(conn net.Conn) {
	err := handshake(conn)
	if err != nil {
		log.Print("error when handshake")
		conn.Close()
		return
	}
	ccfb := cipher.NewCFBEncrypter(c.ClientCipher, iv)
	scfb := cipher.NewCFBDecrypter(c.ServerCipher, iv)
	sConn, err := net.Dial("tcp", c.Server)
	if err != nil {
		log.Print("cannot connect to server")
		conn.Close()
		return
	}
	go HandleStream(sConn, conn, ccfb)
	go HandleStream(conn, sConn, scfb)
}
