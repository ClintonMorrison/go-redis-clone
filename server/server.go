package server

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"go-redis-clone/server/storage"
	"go-redis-clone/server/commands"
)


type Server struct {
	port string
	store *storage.Store
}

func New(port string) *Server {
	server := Server{}
	server.store = storage.NewStore()
	server.port = port

	return &server
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp4", ":" + s.port)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go s.handleRequest(connection)
	}
}

func (s *Server) handleRequest(connection net.Conn) {
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		message := strings.TrimSpace(string(netData))
		if message == "STOP" {
			break
		}

		response := s.processMessage(message)
		connection.Write([]byte(string(response + "\n")))
	}
	connection.Close()
}

func (s *Server) processMessage(message string) string {
	command, err := commands.ParseCommand(message)

	if err == nil {
		return command.Apply(s.store)
	}

	return err.Error()
}


