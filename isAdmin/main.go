package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Admin:", admin())
}

func admin() bool {
	euid := os.Geteuid() // "effective user id", 0 for sudo on Linux, -1 Windows
	if euid == 0 {
		return true
	}
	if euid == -1 {
		f, e := os.Open("\\\\.\\PHYSICALDRIVE0")
		defer f.Close()
		if e == nil { // will have permission error if not elevated on Windows
			return true
		}
	}
	return false
}
