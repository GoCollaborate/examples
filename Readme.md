# GoCollaborate
## What is GoCollaborate?
GoCollaborate is an universal framework for distributed services management that you can easily program with, build extension on, and on top of which you can create your own high performance distributed services like a breeze.
## The Idea Behind
GoCollaborate absorbs the best practice experience from popular distributed services frameworks like[✨Hadoop](https://hadoop.apache.org/), [✨ZooKeeper](https://zookeeper.apache.org/), [✨Dubbo](http://dubbo.io/) and [✨Kite](https://github.com/koding/kite) that helps to ideally resolve the communication and collaboration issues among multiple isolated peer servers.
## Am I Free to Use GoCollaborate?
Yes! Please check out the terms of the BSD License.
## Documents
![alt text](https://github.com/HastingsYoung/GoCollaborate/raw/master/home.png "Docs Home Page")
## Relative Links
Source code: https://github.com/HastingsYoung/GoCollaborate
Examples: https://github.com/HastingsYoung/GoCollaborateExamples
## Quick Start
### Installation
```sh
go get -u github.com/GoCollaborate
```
### Create Project
```sh
mkdir Your_Project_Name
cd Your_Project_Name
mkdir core
touch contact.json
touch main.go
cd ./core
touch simpleTaskHandler.go
```
The project structure now looks something like this:
```
- Your_Project_Name
	- core
		- simpleTaskHandler.go
	- contact.json
	- main.go
```
Configure file `contact.json`:
```js
{
	"agents": [{
		"ip": "localhost",
		"port": 57851,
		"alive": true
	}, {
		"ip": "localhost",
		"port": 57852,
		"alive": true
	}],
	"local": {
		"ip": "localhost",
		"port": 57851,
		"alive": true
	},
	"coordinator": {
		"ip": "",
		"port": 0,
		"alive": false
	},
	"timestamp": 1504182648
}
```
### Entry
```go
import (
	"./core"
	"github.com/GoCollaborate"
)

func main() {
	mp := new(core.SimpleMapper)
	rd := new(core.SimpleReducer)
	collaborate.Set("Function", core.ExampleFunc, "exampleFunc")
	collaborate.Set("Mapper", mp)
	collaborate.Set("Reducer", rd)
	collaborate.Set("Shared", []string{"GET", "POST"}, core.ExampleTaskHandler)
	collaborate.Run()
}
```
### Map-Reduce
```go
import (
	"fmt"
	"github.com/GoCollaborate/server/task"
	"net/http"
)

func ExampleTaskHandler(w http.ResponseWriter, r *http.Request) task.Task {
	return task.Task{task.PERMANENT,
		task.BASE, "exampleFunc", []task.Countable{1, 2, 3, 4},
		[]task.Countable{0},
		task.NewTaskContext(struct{}{})}
}

func ExampleFunc(source []task.Countable,
	result []task.Countable,
	context *task.TaskContext) chan bool {
	out := make(chan bool)
	// deal with passed in request
	go func() {
		fmt.Println("Example Task Executed...")
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
```
### Run
Here we create the entry file and a simple implementation of map-reduce interface, and next we will run with std arguments:
```sh
go run main.go -svrmode=clbt
```
The task is now up and running at:
```
http://localhost:8080/core/ExampleTaskHandler
```