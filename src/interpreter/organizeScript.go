package interpreter

import (
	"pine/src/actions"
	"pine/src/syntax"
	"regexp"
	"strings"
)

func organize(script string) {
	// remove all newline chars using regex
	script = regexp.MustCompile(`\r?\n`).ReplaceAllString(script, " ")

	// remove all spaces
	script = strings.Replace(script, " ", "", -1)

	// list of commands, strings that will later be processed
	var commands []string

	// temporary command where characters will be loaded, and when the whole thing is loaded, reset it
	var command string

	// reading comment, when on, don't record/sort characters
	readingComment := false

	// loop to go over every character in script
	for i := 0; i < len(script); i++ {
		// get currently indexed character
		char := string(script[i])

		// if found beginning of comment
		if char == syntax.Keys["STARTCOMMENT"] {
			readingComment = true // enable reading comment
		} else if char == syntax.Keys["ENDCOMMENT"] {
			readingComment = false // disable reading comment
		}

		// if reading comment is disabled here, the end character of the comment will still be recorded
		// so also check if not reading end comment char

		// if not reading a comment or end comment character
		if !readingComment && char != syntax.Keys["ENDCOMMENT"] {
			// if hit ENDCOMMAND char
			if char == syntax.Keys["ENDCOMMAND"] {
				// append new object to list of strings, which is commands
				commands = append(commands, command)
				// reset command back to empty string
				command = ""
			} else {
				// otherwise add character to command
				command += char
			}
		}
	}
	// run commands by passing them to the actions module and letting it handle it now
	actions.RunActions(commands)
}
