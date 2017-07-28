package main

import (
	"encoding/json"
	"fmt"
	"os/user"
)

func main() {
	me, err := user.Current()
	if err != nil {
		fmt.Println("error getting current user:", err)
	}
	b, err := json.MarshalIndent(me, "", "  ")
	if err != nil {
		fmt.Println("marshal user error:", err)
	}
	fmt.Println("User:")
	fmt.Println(string(b))
	fmt.Println("Group:")
	fmt.Println(me.GroupIds()) //GroupIds not implemented on windows
}
