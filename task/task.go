package task

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/docker/go-connections/nat"
	"time"
)

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID uuid.UUID
	Name string
	State Stat
	Image string
	Memory int
	Disk int
	ExposedPorts nat.PortSet
	PortBindings map[string][string]
	RestartPolicy string
	StartTime time.Time
	FinishTime time.Time
}

type TaskEvent struct {
	ID uuid.UUID
	State State
	TimeStamp time.Time
	Task Task
}

func Task() {
	fmt.Println("Task setup")
}