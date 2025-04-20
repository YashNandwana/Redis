package server

import (
	"fmt"
	"net"
)

type Handler interface {
	Handle(req Request) (Response, error)
}

type Server interface {
	Serve() error
}

type tcpServer struct {
	addr    string
	parser  Parser
	handler Handler
}

func NewServer(addr string, parser Parser, handler Handler) Server {
	return &tcpServer{addr: addr, parser: parser, handler: handler}
}

func (s *tcpServer) Serve() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to bind to %s: %w", s.addr, err)
	}
	defer ln.Close()
	fmt.Println("Server listening on", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *tcpServer) handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		req, err := s.parser.ReadRequest(conn)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		resp, err := s.handler.Handle(req)
		if err != nil {
			fmt.Println("Handle error:", err)
			return
		}

		if err := s.parser.WriteResponse(conn, resp); err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}
