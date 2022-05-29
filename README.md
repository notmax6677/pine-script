# pine-script

pine-script is a language developed in Go, for Go, so far it is kept pretty minimal, but it's open to everyone to build off and modify!

## overview
this is a small language that allows you to create objects and variables in pine-script, then run an interpreter to convert the objects to Go code, which you can then use freely to your liking

the language currently features only two object types (strings and numbers) and support for comments

due to it's small scale and somewhat simple self-documented code, it's pretty easy to modify the syntax, whether changing keys, adding features, or removing some you don't like

## how to use
for starters, you would need to have a `.pine` file in your project folder, once you have this module installed, you may use it to run `.pine` scripts

e.g:
```go
import "pine/src/interpreter" // import the interpreter

func main() {
  interpreter.RunScript("./script.pine") // run your script
  
  objects := interpreter.ExportObjects() // get objects that were created with your pine script
}```
