package main

import (
	"./core"
	"github.com/GoCollaborate"
	"github.com/GoCollaborate/server/chainer"
)

func main() {
	// example 01
	simplmp := new(core.SimpleMapper)
	chainmp := chainer.DefaultChainMapper(0)
	chainmp.Append(simplmp, simplmp, simplmp)

	rd := new(core.SimpleReducer)

	collaborate.Set("Function", core.ExampleFunc, "exampleFunc")
	collaborate.Set("Mapper", simplmp, "core.ExampleTask.SimpleMapper")
	collaborate.Set("Mapper", chainmp, "core.ExampleTask.ChainMapper")
	collaborate.Set("Reducer", rd, "core.ExampleTask.SimpleReducer")
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleJobHandler01)

	// example 02
	advmp := new(core.AdvancedMapper)
	advrd := new(core.AdvancedReducer)
	pipmp := chainer.DefaultPipelineMapper()
	pipmp.Set(advmp, advrd, advmp)

	collaborate.Set("Mapper", pipmp, "core.ExampleTask.PipelineMapper")
	collaborate.Set("Reducer", advrd, "core.ExampleTask.AdvancedReducer")
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleJobHandler02)
	collaborate.Run()
}
