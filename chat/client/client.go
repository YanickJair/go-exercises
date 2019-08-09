package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting chat...")
	fmt.Println("What's your username?")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s, connecting to chat system...\n", name)
	// Dial makes a "call" and try to connect to the chat server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("Could not connect to chat system", err)
	}
	fmt.Println("Connected to chat system")
	name += ":"
	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name+msg+"\n")
	}
}
