package interpreter

import "pine/src/actions"

// list of objects
type objectList struct {
	Numbers []actions.NumberVariable
	Strings []actions.StringVariable
}

// exports all created objects
func ExportObjects() objectList {
	// return objects
	return objectList{Numbers: actions.Numbers, Strings: actions.Strings}
}
