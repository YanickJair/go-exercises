package appliance

// Stove - microwave type
type Stove struct {
	typeName string
}

// Start - method from the Appliance "Father"
func (fr *Stove) Start() {
	fr.typeName = "Microwave"
}

// GetPurpose - implemented from Appliance father
func (fr *Stove) GetPurpose() string {
	return "I am a " + fr.typeName + " I warm stuff up!"
}
