package core

import (
	"fmt"
	"github.com/GoCollaborate/artifacts/task"
	"github.com/GoCollaborate/wrappers/taskHelper"
	"net/http"
)

func ExampleJobHandler(w http.ResponseWriter, r *http.Request) *task.Job {
	job := task.MakeJob()
	job.Tasks(&task.Task{task.SHORT,
		task.BASE, "exampleFunc",
		task.Collection{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
		task.Collection{0},
		task.NewTaskContext(struct{}{}), 0})
	job.Stacks("core.ExampleTask.Mapper", "core.ExampleTask.Reducer")
	return job
}

func ExampleFunc(source *task.Collection,
	result *task.Collection,
	context *task.TaskContext) bool {
	// deal with passed in request
	fmt.Println("Example Task Executed...")
	var total int
	// the function will calculate the sum of source data
	for _, n := range *source {
		total += n.(int)
	}
	result.Append(total)
	return true
}

type SimpleMapper int

func (m *SimpleMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
	// slice the data source of the map into 3 separate segments
	return taskHelper.Slice(inmaps, 3), nil
}

type SimpleReducer int

func (r *SimpleReducer) Reduce(maps map[int]*task.Task) (map[int]*task.Task, error) {
	var sum int
	for _, s := range maps {
		for _, r := range (*s).Result {
			sum += r.(int)
		}
	}
	fmt.Printf("The sum of numbers is: %v \n", sum)
	fmt.Printf("The task set is: %v", maps)
	return maps, nil
}
