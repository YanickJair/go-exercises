package utils

// Set - create a map set struct
type Set map[string]struct{}

// MakeSet -create a new empty set with keys
func MakeSet() Set {
	set := make(Set)
	set["item"] = struct{}{}
	set["item1"] = struct{}{}
	return set
}
