package core

import (
	"fmt"
	"github.com/GoCollaborate/src/artifacts/task"
	"github.com/GoCollaborate/src/wrappers/taskHelper"
	"net/http"
	"time"
)

func ExampleJobHandler(w http.ResponseWriter, r *http.Request, bg *task.Background) {
	job := task.MakeJob()
	delay := 5 * time.Second

	go func() {
		sourceData := task.NewCollection()

		// imagine you have a huge data set to load in
		for i := 0; i < 1000000; i++ {
			sourceData.Append(i)
		}

		// and it takes ages
		<-time.After(delay)

		job.Tasks(
			&task.Task{
				task.SHORT,
				task.BASE,
				"exampleFunc",
				*sourceData,
				task.Collection{0},
				task.NewTaskContext(struct{}{}),
				0,
			},
		)

		job.Stacks("core.ExampleTask.Mapper", "core.ExampleTask.Reducer")

		bg.Mount(job)
	}()

	// this will enable you to return an HTTP call in advance to the loading of data
	w.Write(
		[]byte(
			fmt.Sprintf(
				"Job[%v] successfully added, you'll be seeing it after %v seconds",
				job.Id(),
				delay.Seconds(),
			),
		),
	)
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
