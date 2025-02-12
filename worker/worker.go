package worker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/golang-collections/collections/queue"
	"github.com/DigvijayWagh22/Cube/task"
)

type Worker struct {
	Name string
	Queue queue.Queue
	Db map[uuid.UUID]task.Task
	TaskCount int
}

func(w *Worker) CollectStats() {
	fmt.Println("I'll collect the stats")
}
func(w *Worker) RunTask() {
	fmt.Println("I'll start or stop a task")
}
func(w *Worker) StartTask() {
	fmt.Println("I'll start a task")
}