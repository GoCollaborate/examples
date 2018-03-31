package core

import (
	"fmt"
	"github.com/GoCollaborate/src/artifacts/task"
	"github.com/GoCollaborate/src/helpers/taskHelper"
	"net/http"
)

func ExampleJobHandler02(w http.ResponseWriter, r *http.Request, bg *task.Background) {
	job := task.MakeJob()
	job.Tasks(
		&task.Task{
			task.SHORT,
			task.BASE, "exampleFunc",
			task.Collection{
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			},
			task.Collection{0},
			// stage 0
			task.NewTaskContext(struct{}{}), 0},
		&task.Task{
			task.SHORT,
			task.BASE, "exampleFunc",
			task.Collection{
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			},
			task.Collection{0},
			// stage 0
			task.NewTaskContext(struct{}{}), 0},
		&task.Task{task.SHORT,
			task.BASE, "exampleFunc",
			task.Collection{
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			},
			task.Collection{0},
			// stage 1
			task.NewTaskContext(struct{}{}), 1},
		&task.Task{task.SHORT,
			task.BASE, "exampleFunc",
			task.Collection{
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
				1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4,
			},
			task.Collection{0},
			// stage 1
			task.NewTaskContext(struct{}{}), 1})
	// map twice, reduce once at stage 0
	job.Stacks("core.ExampleTask.AdvancedMapper", "core.ExampleTask.AdvancedMapper", "core.ExampleTask.AdvancedReducer")
	// map once, reduce twice at stage 1
	job.Stacks("core.ExampleTask.AdvancedMapper", "core.ExampleTask.AdvancedReducer", "core.ExampleTask.AdvancedReducer")

	bg.Mount(job)

	w.Write([]byte(fmt.Sprintf("The Job %v Has Been Executed", job.Id())))
}

type AdvancedMapper int

func (m *AdvancedMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
	return taskHelper.Slice(inmaps, 3), nil
}

type AdvancedReducer int

func (r *AdvancedReducer) Reduce(maps map[int]*task.Task) (map[int]*task.Task, error) {
	var (
		sum       int
		sortedSet = make([]*task.Task, 0)
	)

	// return the sorted keys
	for _, k := range taskHelper.Keys(maps) {
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
