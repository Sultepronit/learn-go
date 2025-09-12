package main

import (
	"os"
	"text/template"
)

// Go offers built-in support for creating dynamic content or showing customized output to the user with the text/template package. A sibling package named html/template provides the same API but has additional security features and should be used for generating HTML

func main() {
	// We can create a new template and parse its body from a string. Templates are a mix of static text and “actions” enclosed in {{...}} that are used to dynamically insert content.
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Alternatively, we can use the template.Must function to panic in case Parse returns an error. This is especially useful for templates initialized in the global scope.
	t1 = template.Must(t1.Parse("value: {{.}}\n"))

	// By “executing” the template we generate its text with specific values for its actions. The {{.}} action is replaced by the value passed as a parameter to Execute.
	t1.Execute(os.Stdout, "some text") // value: some text
	t1.Execute(os.Stdout, 777)         // value: 777
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	}) // value: [Go Rust C++ C#]git
}
