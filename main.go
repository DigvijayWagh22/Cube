package main

import (
	"fmt"
	"github.com/DigvijayWagh22/Cube/manager"
	"github.com/DigvijayWagh22/Cube/node"
	"github.com/DigvijayWagh22/Cube/scheduler"
	// "github.com/DigvijayWagh22/Cube/task"
	// "github.com/DigvijayWagh22/Cube/worker"
)

func main() {
	fmt.Println("Orchestrator in golang")
	manager.Manager()
	scheduler.Scheduler()
	node.Node()
	
}