package appliance

// Fridge - type definition
type Fridge struct {
	typeName string
}

// Start - method from the Appliance "Father"
func (fr *Fridge) Start() {
	fr.typeName = "Fridge"
}

// GetPurpose - implemented from Appliance father
func (fr *Fridge) GetPurpose() string {
	return "I am a " + fr.typeName + " I cool stuff down!"
}
