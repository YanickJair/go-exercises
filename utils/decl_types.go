package utils

// CMember - a reference to the crew menbers
type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearence int
}

// New - create a new crew
func New() CMember {
	cM := CMember{
		name:      "Giovanni Gabriel",
		age:       06,
		rank:      "Captain",
		address:   "Praia",
		clearence: 10,
	}
	return cM
}

// Map - use the map keyword to creat a map of users
func Map() map[string]CMember {
	gabriel := New()
	cms := map[string]CMember{
		"gabriel": gabriel,
		"yannick": CMember{
			name:      "Yanick Andrade",
			address:   "Praia",
			age:       26,
			clearence: 9,
			rank:      "Soldier",
		},
	}
	return cms
}
