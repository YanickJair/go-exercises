package utils

import (
	"math/rand"
	"time"
)

/*
* Select Statement: Allows code to wait on multiple channels at the same time
* Select blocks until one channel is ready
* If multiple channels are ready, select picks one at random
 */

var scMapping = map[string]int{
	"Yanick":  26,
	"Djamila": 27,
	"Gabriel": 06,
}

// FindError - implementing our error struct
type FindError struct {
	Name, Server, Msg string
}

func (e FindError) Error() string {
	return e.Msg
}

// FindOne - return random user in our map
func FindOne(name, server string, c chan int) {
	time.Sleep(time.Duration(rand.Intn(60)) * time.Millisecond)

	c <- scMapping[name]
}

// FindSch - implements the Error interface
func FindSch(name, server string) (int, error) {
	time.Sleep(time.Duration(rand.Intn(60)) * time.Millisecond)

	if v, ok := scMapping[name]; ok {
		return v, nil
	}
	return -1, FindError{
		name,
		server,
		"Crew member not found",
	}
}

/*
rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "Gabriel"

	go utils.FindOne(name, "Server 1", c1)
	go utils.FindOne(name, "Server 2", c2)

	select {
	case sc := <-c1:
		fmt.Println(name, " has a security clearence of ", sc, " found in server 1")
	case sc := <-c2:
		fmt.Println(name, " has a security clearence of ", sc, " found in server 2")
	case <-time.After(2 * time.Minute):
		fmt.Println("Search time out!")
	}

	find, err := utils.FindSch("Giovanni", "Server 1")

	if err != nil {
		fmt.Println(err)
		// Assert the type of our error
		if v, ok := err.(utils.FindError); ok {
			fmt.Println("\t- Server name: ", v.Server)
			fmt.Println("\t- Crew member name: ", v.Name)
		}
		return
	}
	fmt.Println(find)
*/
