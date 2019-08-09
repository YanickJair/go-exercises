package chat

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pattern "github.com/YanickJair/intro/singleton-pattern"
)

var logger = pattern.GetInstance()

// Run - Start Chat
func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	if err != nil {
		logger.Println("Error connection to chat client", err)
		return err
	}

	r := CreateRoom("AppChat")
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		fmt.Println("Closing connection...")
		close(r.quit)
		if r.ClCount() > 0 {
			<-r.msgChan
		}
		os.Exit(0)
	}()

	for {
		// Accept a incoming connection
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error connection to chat client", err)
			break
		}
		go handleConnection(r, conn)
	}
	return err
}

//* handle connection from clients and add as a new client
func handleConnection(r *room, c net.Conn) {
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
