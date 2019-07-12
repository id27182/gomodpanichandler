package main

import (
	"fmt"
	"github.com/id27182/gomodpanichandler/anothermodule"
	"github.com/id27182/gomodpanichandler/module"
	"runtime/debug"
	"strings"
)

func main()  {
	defer handlePanicFromModule("module")

	// panic from this module will be handled.
	// comment and rerun to see that panics from anothermodule will not be handled
	module.CausePanic()

	// panic from this module will not be handled
	anothermodule.CausePanic()
}


func handlePanicFromModule(moduleName string) {
	if e := recover(); e != nil {
		// getting a call stack
		stackBytes := debug.Stack()
		stackString := string(stackBytes)

		// dummy way to check if call stack contains module name just to show that it's possible
		if strings.Contains(stackString, fmt.Sprintf("/%s/", moduleName)) {
			fmt.Printf("handling panic in %s", moduleName)
		} else {
			panic(e)
		}
	}
}
