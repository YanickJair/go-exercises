package chat

import (
	"fmt"
	"io"
	"sync"
)

type room struct {
	name    string
	msgChan chan string
	client  map[chan<- string]struct{}
	quit    chan struct{}
	*sync.RWMutex
}

// CreateRoom - initialize a new chat room
func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		msgChan: make(chan string),
		RWMutex: new(sync.RWMutex),
		client:  make(map[chan<- string]struct{}),
		quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

// Run - start the room
func (r *room) Run() {
	logger.Println("Starting chat room", r.name)
	// implement the broadcast sending in a goroutine
	go func() {
		for msg := range r.msgChan {
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
	for wc := range r.client {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}

// AddClient - add a new client to the room
func (r *room) AddClient(c io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.msgChan, c, r.quit)
	r.client[wc] = struct{}{}
	r.Unlock()

	// remove client when done is signalled
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

// ClCount - nº of clients
func (r *room) ClCount() int {
	return len(r.client)
}

// RemoveClient - when client is done, remove from the room
func (r *room) RemoveClient(wc chan<- string) {
	logger.Println("Removing the client...")
	r.Lock()
	//close channel
	close(wc)
	// remove client from the map
	delete(r.client, wc)
	r.Unlock()

	select {
	case <-r.quit:
		if len(r.client) == 0 {
			close(r.msgChan)
		}
	default:
	}
}
