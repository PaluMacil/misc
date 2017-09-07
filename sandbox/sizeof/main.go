package main

import "fmt"
import "unsafe"

func main() {
	s1 := "foo"
	s2 := "foobar"

	fmt.Printf("s1 size: %T, %d\n", s1, unsafe.Sizeof(s1))
	fmt.Printf("s2 size: %T, %d\n", s2, unsafe.Sizeof(s2))
	fmt.Printf("s1 len: %T, %d\n", s1, len(s1))
	fmt.Printf("s2 len: %T, %d\n\n", s2, len(s2))

	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64

	fmt.Printf("size: %T, %d\n", i, unsafe.Sizeof(i))
	fmt.Printf("size: %T, %d\n", i8, unsafe.Sizeof(i8))
	fmt.Printf("size: %T, %d\n", i16, unsafe.Sizeof(i16))
	fmt.Printf("size: %T, %d\n", i32, unsafe.Sizeof(i32))
	fmt.Printf("size: %T, %d\n", i64, unsafe.Sizeof(i64))
	fmt.Printf("size: %T, %d\n", f32, unsafe.Sizeof(f32))
	fmt.Printf("size: %T, %d\n\n", f64, unsafe.Sizeof(f64))

	var b bool
	fmt.Printf("size: %T, %d\n\n", b, unsafe.Sizeof(b))
}
