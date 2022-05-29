package error

import "fmt"

// runs an error message with the given type
func RunError(errType string) {
	// switch statement to iterate through matching types
	switch errType {
	// no file extension
	case "NoFileExtension":
		fmt.Println("None or improper file extension found in given script!")
	// insufficient amount of parameters
	case "InsufficientParams":
		fmt.Println("SYNTAX ERROR: Insufficient amount of parameters!")
	// invalid type given during creation
	case "InvalidTypeCreation":
		fmt.Println("SYNTAX ERROR: Invalid type given during variable creation!")
	}
}
