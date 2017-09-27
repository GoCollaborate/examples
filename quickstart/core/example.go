package core

import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"net/http"
)

func ExampleJobHandler(w http.ResponseWriter, r *http.Request) *task.Job {
	job := task.MakeJob()
	job.Tasks(&task.Task{task.SHORT,
		task.BASE, "exampleFunc",
		[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
		[]task.Countable{0},
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
		var total int
		for _, n := range *source {
			total += n.(int)
		}
		*result = append(*result, total)
		out <- true
	}()
	return out
}

type SimpleMapper int

func (m *SimpleMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
	var (
		s1      []task.Countable
		s2      []task.Countable
		s3      []task.Countable
		s4      []task.Countable
		s5      []task.Countable
		s6      []task.Countable
		gap     = len(inmaps)
		outmaps = make(map[int]*task.Task)
	)
	for k, t := range inmaps {
		var (
			sgap = len(t.Source)
		)
		s1 = t.Source[:sgap/3]
		s2 = t.Source[sgap/3 : sgap*2/3]
		s3 = t.Source[sgap*2/3:]
		s4 = t.Result
		s5 = t.Result
		s6 = t.Result

		outmaps[(k+1)*gap] = &task.Task{t.Type, t.Priority, t.Consumable, s1, s4, t.Context, t.Stage}
		outmaps[(k+1)*gap+1] = &task.Task{t.Type, t.Priority, t.Consumable, s2, s5, t.Context, t.Stage}
		outmaps[(k+1)*gap+2] = &task.Task{t.Type, t.Priority, t.Consumable, s3, s6, t.Context, t.Stage}
	}

	return outmaps, nil
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
