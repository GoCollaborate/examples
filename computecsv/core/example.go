package core

import (
	"fmt"
	"github.com/GoCollaborate/src/artifacts/task"
	"github.com/GoCollaborate/src/wrappers/ioHelper"
	"github.com/GoCollaborate/src/wrappers/taskHelper"
	"net/http"
)

// The following example shows how to caculate a total bill from a csv file
func ExampleJobHandler(w http.ResponseWriter, r *http.Request) *task.Job {
	var (
		job  = task.MakeJob()
		path = "./data.csv"
		raw  = []struct {
			Balance float64
		}{}
		source = task.Collection{}
	)

	ioHelper.FromPath(path).NewCSVOperator().Fill(&raw)

	for _, r := range raw {
		source.Append(task.Countable(r.Balance))
	}

	job.Tasks(&task.Task{task.SHORT,
		task.BASE, "exampleFunc",
		source,
		task.Collection{},
		task.NewTaskContext(struct{}{}), 0})
	job.Stacks("core.ExampleTask.Mapper", "core.ExampleTask.Reducer")
	return job
}

func ExampleFunc(source *task.Collection,
	result *task.Collection,
	context *task.TaskContext) bool {
	// deal with passed in request
	fmt.Println("Example Task Executed...")
	var total float64
	for _, n := range *source {
		total += n.(float64)
	}
	*result = append(*result, total)
	return true
}

type SimpleMapper int

func (m *SimpleMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
	// slice the data source of the map into 3 separate segments
	return taskHelper.Slice(inmaps, 3), nil
}

type SimpleReducer int

func (r *SimpleReducer) Reduce(maps map[int]*task.Task) (map[int]*task.Task, error) {
	var sum float64
	for _, s := range maps {
		for _, r := range (*s).Result {
			sum += r.(float64)
		}
	}
	fmt.Printf("The sum of balance is: %v \n", sum)
	return maps, nil
}
