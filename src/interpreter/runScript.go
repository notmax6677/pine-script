package interpreter

import (
	"fmt"
	"io/ioutil"
	"pine/src/error"
	"pine/src/syntax"
)

// run the given pine script
func RunScript(path string) {

	// get file specified in path
	file, err := ioutil.ReadFile(path)

	// if error isn't nil
	if err != nil {
		// print error
		fmt.Print(err)
	} else if path[len(path)-5:] != syntax.Keys["FILEXT"] { // get last 5 characters of name of script and see if it matches with extension
		error.RunError("NoFileExtension")
	} else { // otherwise
		organize(string(file)) // organize file
	}
}
