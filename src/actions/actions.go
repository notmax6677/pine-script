package actions

import (
	"fmt"
	"pine/src/error"
	"pine/src/syntax"
	"strconv"
)

// run given commands and actions
func RunActions(actions []string) {
	// where all of the actions will be loaded when they are sorted through
	var organizedActions [][]string

	// iterate thru list of actions
	for i := 0; i < len(actions); i++ {
		// array/slice of sections
		var sections []string

		// section where characters will be added to
		var section string

		// iterate through characters of action
		for c := 0; c < len(actions[i]); c++ {
			// create character
			char := string(actions[i][c])

			// if character is equivalent to the split key character
			if char == syntax.Keys["COMMANDSPLIT"] {
				// append what has been gathered of the actions
				sections = append(sections, section)
				// reset section variable back to empty string
				section = ""
			} else { // otherwise
				// continue normally and add character to current section
				section += char
			}
		}

		// append what has been gathered of the actions
		sections = append(sections, section)

		// add sections of action to all organized actions
		organizedActions = append(organizedActions, sections)
	}

	// iterate through the newly collected data of actions
	for i := 0; i < len(organizedActions); i++ {
		// attempt to execute each one
		executeAction(organizedActions[i])
	}
}

// executes a given action, which is passed on in the form of a slice/array of indexed strings
func executeAction(action []string) {
	// switch through the first item of the action to see the type of the action
	switch action[0] {
	// create variable action
	case syntax.Keys["CREATEVAR"]:
		// if length of action is 4
		if len(action) == 4 {
			// if variable type is string
			if action[1] == syntax.Keys["STRINGTYPE"] {
				// call create string function and pass value
				createString(action[2], action[3])
			} else if action[1] == syntax.Keys["NUMTYPE"] { // if variable type is number
				// attempt to convert value from string
				convertedValue, err := strconv.ParseFloat(action[3], 64)

				// if encountered no errors during conversion
				if err == nil {
					createNumber(action[2], convertedValue) // create number with created value
				} else { // if an error has been found
					fmt.Println(err) // print the error
				}
			} else {
				// run variable creation error
				error.RunError("InvalidTypeCreation")
			}
		} else { // otherwise
			error.RunError("InsufficientParams") // run variable creation error
		}
	}
}
