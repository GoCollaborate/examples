package core

import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"net/http"
)

func TaskAHandler(w http.ResponseWriter, r *http.Request) task.Task {
	return task.Task{task.PERMANENT,
		task.BASE, "funcA", []task.Countable{1, 2, 3, 4},
		[]task.Countable{0},
		task.NewTaskContext(struct{}{})}
}

func FuncA(source []task.Countable,
	result []task.Countable,
	context *task.TaskContext) chan bool {
	out := make(chan bool)
	// deal with passed in request
	go func() {
		fmt.Println("Task A Executed...")
		out <- true
	}()
	return out
}

type SimpleMapper int

func (m *SimpleMapper) Map(t *task.Task) (map[int64]task.Task, error) {
	maps := make(map[int64]task.Task)
	s1 := t.Source[0:1]
	s2 := t.Source[1:2]
	s3 := t.Source[2:3]
	s4 := t.Source[3:4]
	s5 := t.Result
	s6 := t.Result
	s7 := t.Result
	s8 := t.Result
	maps[int64(0)] = task.Task{t.Type, t.Priority, t.Consumable, s1, s5, t.Context}
	maps[int64(1)] = task.Task{t.Type, t.Priority, t.Consumable, s2, s6, t.Context}
	maps[int64(2)] = task.Task{t.Type, t.Priority, t.Consumable, s3, s7, t.Context}
	maps[int64(3)] = task.Task{t.Type, t.Priority, t.Consumable, s4, s8, t.Context}
	return maps, nil
}

type SimpleReducer int

func (r *SimpleReducer) Reduce(sources map[int64]task.Task, result *task.Task) error {
	rs := *result
	for _, s := range sources {
		rs.Result = append(rs.Result, s.Result...)
	}
	return nil
}
