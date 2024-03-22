package main

import ("fmt"
       
	   "log"
	   "github.com/cerbos/cerbos-sdk-go/cerbos"
)

func main() {
  	// Initialize Cerbos client with adress of the PDP
	cerbosClient, err := client.NewClient("http://cerbos:3592")
	if err != nil {
		log.Fatalf("Error initializing Cerbos client: %v", err)
	}
}
