package core

import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"net/http"
)

func ExampleJobHandler02(w http.ResponseWriter, r *http.Request) *task.Job {
	job := task.MakeJob()
	job.Tasks(&task.Task{task.PERMANENT,
		task.BASE, "exampleFunc",
		[]task.Countable{1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 15, 16,
			17, 18, 19, 20, 21},
		[]task.Countable{0},
		task.NewTaskContext(struct{}{}), "core.ExampleTask.PipelineMapper", "core.ExampleTask.AdvancedReducer", 0})
	return job
}

type AdvancedReducer int

func (r *AdvancedReducer) Reduce(source map[int64]*task.Task, result *task.Task) error {
	rs := *result
	var sum int
	for _, s := range source {
		sum += len((*s).Result)
	}
	rs.Source = append(rs.Result, sum)
	fmt.Printf("The number of 3s is: %v \n", sum)
	fmt.Printf("The task result set is: %v", rs)
	return nil
}

type AdvancedMapper int

// this is a mapper designed to filter the multiples of 3
func (m *AdvancedMapper) Map(t *task.Task) (map[int64]*task.Task, error) {
	maps := make(map[int64]*task.Task)
	for i, s := range t.Source {
		if i%3 != 0 {
			continue
		}

		maps[int64(i)] = &task.Task{t.Type, t.Priority, t.Consumable, []task.Countable{s}, []task.Countable{s}, t.Context, "core.ExampleTask.AdvancedMapper", t.Reducer, t.Stage}
	}
	return maps, nil
}
