package main

import (
	"go-redis-clone/server"
	"go-redis-clone/client"
	"fmt"
)

func main() {
	host := "localhost"
	port := "3000"
	address := fmt.Sprintf("%s:%s", host, port)

	s := server.New(port)
	go s.Listen()

	c := client.New(address, true)
	c.Connect()
	c.Send("PING")
	c.Send("SET test 1234")
	c.Send("GET test")
	c.Send("KEYS")
	c.Send("DELETE test")
	c.Send("KEYS")
}
