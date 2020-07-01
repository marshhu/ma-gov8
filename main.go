package main

import (
	"fmt"
    "github.com/lazytiger/go-v8"
)

func main() {
	engine := v8.NewEngine()
	script := engine.Compile([]byte("'Hello ' + 'World!'"), nil, nil)
	context := engine.NewContext(nil)

	context.Scope(func(cs v8.ContextScope) {
		result := script.Run()
		fmt.Println(result.ToString())
	})
}
