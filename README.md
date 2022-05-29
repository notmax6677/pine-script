# pine-script

pine-script is a language developed in Go, for Go, so far it is kept pretty minimal, but it's open to everyone to build off and modify!

## overview
this is a small language that allows you to create objects and variables in pine-script, then run an interpreter to convert the objects to Go code, which you can then use freely to your liking

the language currently features only two object types (strings and numbers) and support for comments

due to it's small scale and somewhat simple self-documented code, it's pretty easy to modify the syntax, whether changing keys, adding features, or removing some you don't like

## how to use
for starters, you would need to have a `.pine` file in your project folder, idk how to make this a proper module so just download it lol

e.g:
```go
import "pine/src/interpreter" // import the interpreter

func main() {
  interpreter.RunScript("./script.pine") // run your script
  
  objects := interpreter.ExportObjects() // get objects that were created with your pine script
}
```

within objects you will find two fields; Numbers & Strings

let's use the example pine script as an example, if i wanted to get its name i would write

`objects.Strings[0].Name`

or if i wanted the year's value, i could do

`objects.Numbers[0].Value`

both strings and numbers have a Name and Value field, but the value is either `string` or `float64` depending on the type

creating a variable is pretty simple;

`var:type:name:value;`

e.g:
```
var:str:MyName:Max;

var:num:MyAge:14;
```

the language filters out all spaces and newlines upon the start of interpretation
so something like
`var:st r:ca r: mer c e d es;`
would convert to
`var:str:car:mercedes;`

if you want to use spaces within the string's value, you may use an alias, named `"SPACECHAR"` in `keys.go`, it defaults to `#`

so

`var:str:car:hello#world`

the value of this variable named car would be `"hello world"`

ps: i don't really care if it doesn't do anything useful, i just wanted to see if i could make something of this sort
