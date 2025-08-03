package main

import (
	"fmt"
	"net"
)

func main() {
	message := "Server working!"
	lisnere, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lisnere.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := lisnere.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
