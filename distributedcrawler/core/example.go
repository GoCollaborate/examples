package core

import (
	"fmt"
	"github.com/GoCollaborate/artifacts/task"
	"github.com/GoCollaborate/wrappers/ioHelper"
	"github.com/GoCollaborate/wrappers/taskHelper"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// The following example shows how to collect HTML contents from the given urls
func ExampleJobHandler(w http.ResponseWriter, r *http.Request) *task.Job {
	var (
		job  = task.MakeJob()
		path = "./data.csv"
		raw  = []struct {
			URL string
		}{}
		source = task.Collection{}
	)

	ioHelper.FromPath(path).NewCSVOperator().Fill(&raw)

	for _, r := range raw {
		source.Append(task.Countable(r.URL))
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
	var text = task.Collection{}

	for _, n := range *source {
		var (
			bytes []byte
			resp  *http.Response
			err   error
		)
		resp, err = http.Get(n.(string))
		if err != nil {
			break
		}
		bytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			break
		}
		text = append(text, task.Countable(string(bytes)))
	}

	*result = append(*result, text...)
	return true
}

type SimpleMapper int

func (m *SimpleMapper) Map(inmaps map[int]*task.Task) (map[int]*task.Task, error) {
	// slice the data source of the map into 3 separate segments
	return taskHelper.Slice(inmaps, 3), nil
}

type SimpleReducer int

func (r *SimpleReducer) Reduce(maps map[int]*task.Task) (map[int]*task.Task, error) {
	var (
		sum  int
		text string = ""
	)
	for _, s := range maps {
		sum += len((*s).Result)
		for _, r := range (*s).Result {
			text += r.(string)
		}
	}
	file, _ := os.Create("./websites.txt")
	io.WriteString(file, text)
	fmt.Printf("The sites visited: %v \n", sum)
	return maps, nil
}
