package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := make([]byte, 4)
	fmt.Println("Hello World", len(data))

	num := uint32(0x12345678)

	// Big-endian (network byte order)
	bigEndian := make([]byte, 4)
	binary.BigEndian.PutUint32(bigEndian, num)
	fmt.Printf("Big-endian:    %x\n", bigEndian) // [12 34 56 78]

	// Little-endian (x86 byte order)
	littleEndian := make([]byte, 4)
	binary.LittleEndian.PutUint32(littleEndian, num)
	fmt.Printf("Little-endian: %x\n", littleEndian) // [78 56 34 12]

	// Reading back with wrong endianness
	wrong := binary.LittleEndian.Uint32(bigEndian)
	fmt.Printf("Wrong read: 0x%x (should be 0x12345678)\n", wrong)
	// Output: Wrong read: 0x78563412Î

}
