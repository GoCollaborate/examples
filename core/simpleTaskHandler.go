package core

import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"net/http"
)

func ExampleTaskHandler(w http.ResponseWriter, r *http.Request) task.Task {
	return task.Task{task.PERMANENT,
		task.BASE, "exampleFunc", []task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
		[]task.Countable{0},
		task.NewTaskContext(struct{}{}), "core.ExampleTaskHandler.Mapper", "core.ExampleTaskHandler.Reducer"}
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

func (m *SimpleMapper) Map(t *task.Task) (map[int64]*task.Task, error) {
	maps := make(map[int64]*task.Task)
	s1 := t.Source[:4]
	s2 := t.Source[4:8]
	s3 := t.Source[8:]
	s4 := t.Result
	s5 := t.Result
	s6 := t.Result
	maps[int64(0)] = &task.Task{t.Type, t.Priority, t.Consumable, s1, s4, t.Context, t.Mapper, t.Reducer}
	maps[int64(1)] = &task.Task{t.Type, t.Priority, t.Consumable, s2, s5, t.Context, t.Mapper, t.Reducer}
	maps[int64(2)] = &task.Task{t.Type, t.Priority, t.Consumable, s3, s6, t.Context, t.Mapper, t.Reducer}
	return maps, nil
}

type SimpleReducer int

func (r *SimpleReducer) Reduce(source map[int64]*task.Task, result *task.Task) error {
	rs := *result
	var sum int
	for _, s := range source {
		for _, r := range (*s).Result {
			sum += r.(int)
		}
	}
	rs.Result[0] = sum
	fmt.Printf("The sum of numbers is: %v \n", sum)
	fmt.Printf("The task result set is: %v", rs)
	return nil
}
