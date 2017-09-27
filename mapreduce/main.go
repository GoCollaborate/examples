package main

import (
	"./core"
	"github.com/GoCollaborate"
)

func main() {
	// example 01
	simplmp := new(core.SimpleMapper)
	simplrd := new(core.SimpleReducer)

	collaborate.Set("Function", core.ExampleFunc, "exampleFunc")
	collaborate.Set("Mapper", simplmp, "core.ExampleTask.SimpleMapper")
	collaborate.Set("Reducer", simplrd, "core.ExampleTask.SimpleReducer")
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleJobHandler01)

	// example 02
	advmp := new(core.AdvancedMapper)
	advrd := new(core.AdvancedReducer)

	collaborate.Set("Mapper", advmp, "core.ExampleTask.AdvancedMapper")
	collaborate.Set("Reducer", advrd, "core.ExampleTask.AdvancedReducer")
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleJobHandler02)
	collaborate.Run()
}
