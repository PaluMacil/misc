package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type T struct {
	F1 [5]byte
	F2 int32
}

// Intel is little-endian
func main() {
	var t1, t2 T
	t1 = T{[5]byte{'a', 'b', 'c', 'd', 'e'}, 1234}
	fmt.Println("t1:", t1)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, &t1)
	if err != nil {
		fmt.Println(err)
	}
	err = binary.Read(buf, binary.BigEndian, &t2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("t2:", t2)
}
