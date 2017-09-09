package core

import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"net/http"
)

func ExampleTaskHandler01(w http.ResponseWriter, r *http.Request) task.Task {
	return task.Task{task.PERMANENT,
		task.BASE, "exampleFunc",
		[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
		[]task.Countable{0},
		task.NewTaskContext(struct{}{}), "core.ExampleTaskHandler.ChainMapper", "core.ExampleTaskHandler.SimpleReducer"}
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
	length := len(t.Source)
	l1 := length / 3
	l2 := length / 3 * 2

	s1 := t.Source[:l1]
	s2 := t.Source[l1:l2]
	s3 := t.Source[l2:]
	s4 := t.Result
	s5 := t.Result
	s6 := t.Result

	index, err := t.Context.Get("index")

	if err != nil {
		return maps, err
	}

	i := index.(int64)

	maps[i] = &task.Task{t.Type, t.Priority, t.Consumable, s1, s4, t.Context, "core.ExampleTaskHandler.SimpleMapper", t.Reducer}
	maps[i+1] = &task.Task{t.Type, t.Priority, t.Consumable, s2, s5, t.Context, "core.ExampleTaskHandler.SimpleMapper", t.Reducer}
	maps[i+2] = &task.Task{t.Type, t.Priority, t.Consumable, s3, s6, t.Context, "core.ExampleTaskHandler.SimpleMapper", t.Reducer}
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
