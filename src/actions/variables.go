package actions

import "pine/src/syntax"

// string variable
type StringVariable struct {
	Name string // name

	Value string // string value
}

// number variable, value is float64 for largest possible space
type NumberVariable struct {
	Name string // name

	Value float64 // float64 value
}

// slice of numbers
var Numbers []NumberVariable

// slice of strings
var Strings []StringVariable

// create number variable
func createNumber(name string, value float64) {
	Numbers = append(Numbers, NumberVariable{Name: name, Value: value})
}

// create string variable
func createString(name string, value string) {
	// modified version of value string that will undergo some changes
	modValue := ""

	// iterate thru value string
	for i := 0; i < len(value); i++ {
		// if character is equal to alias for space character
		if string(value[i]) == syntax.Keys["SPACECHAR"] {
			modValue += " " // add space
		} else { // else
			modValue += string(value[i]) // add normal character
		}
	}

	// note: created variable's value is modified new value
	Strings = append(Strings, StringVariable{Name: name, Value: modValue})
}
