package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"radiant/client"
)

func main() {
	c, err := client.NewClient("localhost:8080") // Connect to server at localhost port 8080.
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	fmt.Println("Connected to server...")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a message to send: ")

	for scanner.Scan() {
		msg := scanner.Text()

		err := c.SendMessage(msg)
		if err != nil {
			log.Fatalf("Could not send message: %v", err)
		}

		fmt.Println("Message sent!")
		fmt.Println("Enter a message to send: ")
	}
}
