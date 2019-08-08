package exclusion

import "sync"

// SafeCounter . impliments the Mutex interface
type SafeCounter struct {
	i int
	sync.Mutex
}

// Incriment - up i+1
func (sc *SafeCounter) Incriment() {
	sc.Lock()
	sc.i += 4
	sc.Unlock()
}

// Decrement - down i-1
func (sc *SafeCounter) Decrement() {
	sc.Lock()
	sc.i--
	sc.Unlock()
}

// Get - return i
func (sc *SafeCounter) Get() int {
	sc.Lock()
	i := sc.i
	sc.Unlock()
	return i
}
