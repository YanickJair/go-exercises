package exclusion

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// MapCounter - impliments RWMutex interface
type MapCounter struct {
	m map[int]int
	sync.RWMutex
}

// New - ...
func New() MapCounter {
	return MapCounter{m: make(map[int]int)}
}

// RunWriter - start runner
func RunWriter(mc *MapCounter, n int) {
	for i := 0; i < n; i++ {
		mc.Lock()
		mc.m[i] = rand.Intn(n) * 10
		mc.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// RunReader - read our map struct and use ReadLock and ReadUnlock
func RunReader(mc *MapCounter, n int) {
	for {
		mc.RLock()
		v := mc.m[rand.Intn(n)]
		mc.RUnlock()
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}
