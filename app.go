package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)


func main() {
	go pingOtherServer()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}


func pingOtherServer() {
	for {
		// Ping the other server
		resp, err := http.Get("https://go-lang-server-testing.onrender.com/ping")
		if err != nil {
			fmt.Printf("Error pinging the other server: %v\n", err)
		} else {
			fmt.Printf("Pinged the other server, status code: %d\n", resp.StatusCode)
			resp.Body.Close()
		}

		// Wait for 15 minutes in a separate goroutine
		time.Sleep(10 * time.Minute)
	}
}
