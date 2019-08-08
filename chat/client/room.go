package client

import (
	"fmt"
	"io"
	"net"
	"sync"

	pattern "github.com/YanickJair/intro/singleton-pattern"
)

var logger = pattern.GetInstance()

type room struct {
	name     string
	msg_chan chan string
	client   map[chan<- string]struct{}
	quit     chan struct{}
	*sync.RWMutex
}

// CreateRoom - initialize a new chat room
func CreateRoom(name string) *room {
	room := &room{
		name:     name,
		msg_chan: make(chan string),
		RWMutex:  new(sync.RWMutex),
		client:   make(map[chan<- string]struct{}),
		quit:     make(chan struct{}),
	}
	r.Run()
	return r
}

// Run - start the room
func (r *room) Run() {
	logger.Println("Starting chat room", r.name)
	// implement the broadcast sending in a goroutine
	go func() {
		for msg := range r.msg_chan {
			r.broadCastMsg(msg)
		}
	}()
}

// broadCastMsg - broadcast the received msg
func (r *room) broadCastMsg(msg string) {
	//! Lock the read function
	r.RLock()

	//! RUnlock guarantee that the unock happens in the and of the function
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	// Receive a message and send it to all the clients connected to the chat room
	for wc, _ := range r.client {
		go func(wc chan<- string) {
			wc <- string
		}(wc)
	}
}

// AddClient - add a new client to the room
func (r *room) AddClient(c *io.ReadWriteSeeker) {
	if v, ok := c.(net.Conn); ok {
	}
}
