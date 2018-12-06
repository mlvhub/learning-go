package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	query := "Our Query"
	respond := make(chan string, 2)

	go googleIt(respond, query)
	go bingIt(respond, query)

	queryResp := <-respond

	fmt.Printf("Sent query:\t\t %s\n", query)
	fmt.Printf("Got Response:\t\t %s\n", queryResp)
}

func googleIt(respond chan<- string, query string) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	respond <- "A Google Response"
}

func bingIt(respond chan<- string, query string) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	respond <- "A Bing Response"
}
