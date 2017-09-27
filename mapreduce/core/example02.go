package core

import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"github.com/GoCollaborate/server/taskutils"
	"net/http"
)

func ExampleJobHandler02(w http.ResponseWriter, r *http.Request) *task.Job {
	job := task.MakeJob()
	job.Tasks(
		&task.Task{task.SHORT,
			task.BASE, "exampleFunc",
			[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
			[]task.Countable{0},
			// stage 0
			task.NewTaskContext(struct{}{}), 0},
		&task.Task{task.SHORT,
			task.BASE, "exampleFunc",
			[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
			[]task.Countable{0},
			// stage 0
			task.NewTaskContext(struct{}{}), 0},
		&task.Task{task.SHORT,
			task.BASE, "exampleFunc",
			[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
			[]task.Countable{0},
			// stage 1
			task.NewTaskContext(struct{}{}), 1},
		&task.Task{task.SHORT,
			task.BASE, "exampleFunc",
			[]task.Countable{1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4},
			[]task.Countable{0},
			// stage 1
			task.NewTaskContext(struct{}{}), 1})
	// map once, reduce once at stage 0
	job.Stacks("core.ExampleTask.AdvancedMapper", "core.ExampleTask.AdvancedReducer")
	// map twice, reduce once at stage 1
	job.Stacks("core.ExampleTask.AdvancedMapper", "core.ExampleTask.AdvancedMapper", "core.ExampleTask.AdvancedReducer")
	return job
}

type AdvancedMapper int

func (m *AdvancedMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
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

type AdvancedReducer int

func (r *AdvancedReducer) Reduce(maps map[int]*task.Task) (map[int]*task.Task, error) {
	var (
		sum       int
		sortedSet = make([]*task.Task, 0)
	)

	// return the sorted keys
	for _, k := range taskutils.Keys(maps) {
		s := maps[k]
		sortedSet = append(sortedSet, s)
		for _, r := range (*s).Result {
			sum += r.(int)
		}
	}
	fmt.Printf("The sum of numbers is: %v \n\n", sum)
	fmt.Printf("The task set is: %v \n\n", maps)
	fmt.Printf("The sorted set is: %v \n\n", sortedSet)
	return maps, nil
}
