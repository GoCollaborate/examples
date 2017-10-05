package main

import (
	"./core"
	"github.com/GoCollaborate"
	"time"
)

func main() {
	mp := new(core.SimpleMapper)
	rd := new(core.SimpleReducer)
	collaborate.Set("Function", core.ExampleFunc, "exampleFunc")
	collaborate.Set("Mapper", mp, "core.ExampleTask.Mapper")
	collaborate.Set("Reducer", rd, "core.ExampleTask.Reducer")
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleJobHandler)

	// this will allow the job service token to be refilled every 500 millisecond, the
	// maximum job request running in parallel is 1
	collaborate.Set("Limit", "/core/ExampleJobHandler", 500*time.Millisecond, 1)
	collaborate.Run()
}
