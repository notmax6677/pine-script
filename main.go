package main

import (
	"fmt"
	"pine/src/interpreter"
)

func main() {
	interpreter.RunScript("./script.pine")

	objects := interpreter.ExportObjects()

	fmt.Println(objects)
}
