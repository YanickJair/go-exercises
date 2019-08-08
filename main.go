package main

import (
	"time"

	"github.com/YanickJair/intro/exclusion"
)

func main() {
	mc := exclusion.New()
	go exclusion.RunWriter(&mc, 20)
	go exclusion.RunReader(&mc, 20)
	go exclusion.RunReader(&mc, 20)
	time.Sleep(15 * time.Second)
}
