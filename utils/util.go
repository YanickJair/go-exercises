package utils

import "fmt"

// SayHello - greet our dear user
func SayHello() {
	fmt.Println("Hello")
}

// Say - user define what to say
func Say(s string, i int) {
	fmt.Println(s, i)
}

// Incrimenter - return a function which incriments values
func Incrimenter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// ChangeX - receive a pointer and change is value in memory
func ChangeX(x *int) {
	*x = 10
}
