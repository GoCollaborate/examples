package main

import (
	"./core"
	"github.com/GoCollaborate"
)

func main() {
	mp := new(core.SimpleMapper)
	rd := new(core.SimpleReducer)
	collaborate.Set("Function", core.ExampleFunc, "exampleFunc")
	collaborate.Set("Mapper", mp, "core.ExampleTask.Mapper")
	collaborate.Set("Reducer", rd, "core.ExampleTask.Reducer")
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleJobHandler)
	collaborate.Run()
}
