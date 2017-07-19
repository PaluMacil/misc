package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now:")
	n := time.Now()
	fmt.Println(n)
	fmt.Println(n.Location())
	fmt.Println("UTC:")
	u := n.UTC()
	fmt.Println(u)
	fmt.Println(u.Location())
	fmt.Println("Diff:")
	fmt.Println(n.Sub(u))
	b, e := u.GobEncode()
	fmt.Println(b, e)
	fmt.Println(u.Unix())
	fmt.Println("CET:")
	loc, e := time.LoadLocation("CET")
	fmt.Println(loc, e)
	c := n.In(loc)
	fmt.Println(c)
}
