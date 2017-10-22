package core

import (
	"fmt"
	"github.com/GoCollaborate/artifacts/task"
	"github.com/GoCollaborate/wrappers/ioHelper"
	"github.com/GoCollaborate/wrappers/taskHelper"
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
		source = []task.Countable{}
	)

	ioHelper.FromPath(path).NewCSVOperator().Fill(&raw)

	for _, r := range raw {
		source = append(source, task.Countable(r.Balance))
	}

	job.Tasks(&task.Task{task.SHORT,
		task.BASE, "exampleFunc",
		source,
		[]task.Countable{},
		task.NewTaskContext(struct{}{}), 0})
	job.Stacks("core.ExampleTask.Mapper", "core.ExampleTask.Reducer")
	return job
}

func ExampleFunc(source *[]task.Countable,
	result *[]task.Countable,
	context *task.TaskContext) chan bool {
	out := make(chan bool)
	// deal with passed in request
	go func() {
		fmt.Println("Example Task Executed...")
		var total float64
		for _, n := range *source {
			total += n.(float64)
		}
		*result = append(*result, total)
		out <- true
	}()
	return out
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
	fmt.Printf("The sum of numbers is: %v \n", sum)
	fmt.Printf("The task set is: %v", maps)
	return maps, nil
}
