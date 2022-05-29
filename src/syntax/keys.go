package syntax

// syntax keywords/characters for the language
var Keys = map[string]string{
	"COMMANDSPLIT": ":", // splits commands apart

	"ENDCOMMAND": ";", // end of a command

	"STARTCOMMENT": "<", // beginning of a comment
	"ENDCOMMENT":   ">", // end comment

	"SPACECHAR": "#", // represents a space, if placed in string, will convert to a space

	// possible commands/actions

	"CREATEVAR":  "var", // creates a variable
	"STRINGTYPE": "str", // string variable type
	"NUMTYPE":    "num", // number variable type

	// file extention
	"FILEXT": ".pine",
}
