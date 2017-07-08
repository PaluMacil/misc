package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://api.github.com/repos/palumacil/misc/releases"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("User-Agent", "PaluMacil")
	res, err := client.Do(req)
	checkErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)
	//TODO: print time till reset of rate limit
	//unixNow := time.Now().UTC().Format(time.UnixDate)
	//resetTime := res.Header["X-RateLimit-Reset"]
	//tillReset := ""

	//time.Time{sec=}
	fmt.Println("GET:\n", string(body))
	fmt.Println("Rate Limit Remaining:", res.Header["X-Ratelimit-Remaining"])
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
