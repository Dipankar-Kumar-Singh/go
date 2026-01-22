package main

import (
	"fmt"
	"net"
)

func main() {
	Client()
}

func Client() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	//data := []byte("hello world")
	data := []byte{05, 23, 104, 101, 108, 108, 111, 32}
	write, err := conn.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", write)
}
