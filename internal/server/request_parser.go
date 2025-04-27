package server

import (
	"bufio"
	"fmt"
	"net"
)

type Parser interface {
	ReadRequest(conn net.Conn) (Request, error)
	WriteResponse(conn net.Conn, resp Response) error
}

type RESPParser struct{}

func NewRESPParser() Parser {
	return &RESPParser{}
}

func (p *RESPParser) ReadRequest(conn net.Conn) (Request, error) {
	reader := bufio.NewReader(conn)
	parts, err := parseRESP(reader)
	if err != nil {
		return Request{}, err
	}
	if len(parts) < 1 {
		return Request{}, fmt.Errorf("empty request")
	}
	return Request{Command: parts[0], Args: parts}, nil
}

func (p *RESPParser) WriteResponse(conn net.Conn, resp Response) error {
	fmt.Println(resp)
	data := generateResponse(resp)
	_, err := conn.Write([]byte(data))
	return err
}
