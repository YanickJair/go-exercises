package utils

import "fmt"

/*
* Channels a way to goroutines comunicate
* Creating: ch := mak(chan <type>)
* Sending value: ch <-X
* Receiving value: X := <-ch
 */

/*
* Buffered channels: can be considered as a buffer of stacked channels
* Declaring: ch := make(chan <type>, <number of buffered channels>), i.e: ch := make(chan int, 100)
* Sender only block when channel is full
* Receiver block when channel is empty
 */

/*
* Range and close
* Range can be used to receive values from a channel inside a loop
* Close is used to indicate the channel has retired
 */

// WaitAndSay - wait and say some
func WaitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}
	c <- true
}

// SayHelloMultipleTimes - testing Range and Close concepts
func SayHelloMultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}

/*
c := make(chan bool)
	go utils.WaitAndSay(c, "Talk to me my friend")
	fmt.Println("I'm waiting...")

	c <- true

	<-c

	ch := make(chan string)
	go utils.SayHelloMultipleTimes(ch, 5)
	for s := range ch {
		fmt.Println(s)
	}
	_, ok := <-ch
	fmt.Println("Channel close?", !ok)
*/
