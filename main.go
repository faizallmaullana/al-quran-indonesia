package main

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var server int

func main() {
	ArgunemtHandler()

	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("Could not open log file:", err)
		return
	}
	defer file.Close()
	log.SetOutput(file)

	var wg sync.WaitGroup

	numServers := server

	for i := 0; i < numServers; i++ {
		wg.Add(1)
		go startServer(&wg)
	}

	// Load balancer setup
	lb := createLoadBalancer()
	go func() {
		if err := lb.Run(":3000"); err != nil {
			log.Println("Load balancer failed:", err)
		}
	}()

	wg.Wait() // Wait for all servers to finish (they won't, this is just for completion)
}

func ArgunemtHandler() {
	if len(os.Args) <= 1 {
		panic("Please fill the numserver\n./main [num_server]")
	}

	numservers := os.Args[1]

	num, err := strconv.Atoi(numservers)
	if err != nil {
		panic(err)
	}

	server = num
}
