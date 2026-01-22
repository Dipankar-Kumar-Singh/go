package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	Server()
}

func Server() {
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer func(ln net.Listener) {
		err := ln.Close()
		if err != nil {
			fmt.Println("Close server:", err)
		}
		fmt.Println("Closed the server!")
	}(ln)

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				//fmt.Println("Client disconnected")
				//break // Exit the loop normally
				continue
			}
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		fmt.Println("Raw buffer:", buf[:n])
	}
}
