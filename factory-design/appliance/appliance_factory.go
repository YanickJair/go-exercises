package appliance

import "errors"

// Appliance - the "father" interface in our design
type Appliance interface {
	Start()
	GetPurpose() string
}

// Applicances types
const (
	STOVE = iota
	FRIDGE
)

// CreateAppliance - Create a new Applicane const
func CreateAppliance(mType int) (Appliance, error) {
	switch mType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}

/*
	fmt.Println("Enter preffered appliance type")
	fmt.Println("0: Stove")
	fmt.Println("1: Fridge")

	var appType int
	fmt.Scan(&appType)

	myAppliance, err := appliance.CreateAppliance(appType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	}
*/
