package task_manager

import (
	"container/heap"
	"os/exec"
	"sync"
	"fmt"
	"time"
)

type TaskManager struct {
	taskQueue *Tasks
	sync.Mutex

	taskMap  map[string]*Task
	ExitChan chan struct{}
	WG       sync.WaitGroup
}

func NewTaskManager() (tm *TaskManager) {
	tm = &TaskManager{
		taskQueue: &Tasks{},
		taskMap:   make(map[string]*Task),
		ExitChan:  make(chan struct{}),
		WG:        sync.WaitGroup{},
	}
	heap.Init(tm.taskQueue)

	tm.WG.Add(1)
	go tm.Run()

	return
}

func (tm *TaskManager) AddTask(task *Task) error {
	if _, ok := tm.taskMap[task.ID]; ok {
		return fmt.Errorf(fmt.Sprintf("The task %s already exists.", task.ID))
	}

	tm.Lock()
	defer tm.Unlock()
	tm.taskMap[task.ID] = task
	heap.Push(tm.taskQueue, task)

	return nil
}

func (tm *TaskManager) DelTask(id string) error {
	tm.Lock()
	defer tm.Unlock()

	task, ok := tm.taskMap[id]
	if !ok {
		return fmt.Errorf(fmt.Sprintf("The task %s is not found.", id))
	}

	task.IsActive = false
	delete(tm.taskMap, task.ID)

	return nil
}

func (tm *TaskManager) Run() {
	ticker := time.Tick(100 * time.Microsecond)

	for {
		select {
		case <-ticker:
			if tm.taskQueue.Len() > 0 {
				tm.Lock()
				for {
					if tm.taskQueue.Len() < 1 {
						break
					}
					task := heap.Pop(tm.taskQueue).(*Task)
					if !task.IsActive {
						continue
					}
					if task.NextTime > time.Now().UnixNano() {
						heap.Push(tm.taskQueue, task)
						break
					}

					tm.WG.Add(1)
					go func() {
						buf, err := exec.Command(task.Cmd, task.Args...).Output()
						if err != nil {
							fmt.Println(err.Error())
						}
						fmt.Println(string(buf))
						tm.WG.Done()
					}()

					task.NextTime += task.Interval
					heap.Push(tm.taskQueue, task)
				}
				tm.Unlock()
			}
		case <-tm.ExitChan:
			goto Finish
		}
	}
Finish:
	tm.WG.Done()
}
