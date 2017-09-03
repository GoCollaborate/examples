package main

import (
	"./core"
	"github.com/GoCollaborate"
)

func main() {
	mp := new(core.SimpleMapper)
	rd := new(core.SimpleReducer)
	collaborate.Set("Function", core.ExampleFunc, "exampleFunc")
	collaborate.Set("Mapper", mp)
	collaborate.Set("Reducer", rd)
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleTaskHandler)
	collaborate.Run()
}
