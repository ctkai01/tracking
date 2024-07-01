package models

import (
	"fmt"
	
	"fyne.io/fyne/v2/data/binding"
)


type Task struct {
	Id int
	ProjectName string
	TaskName string
	Status int
	WorkingTime int
}

func NewTask(id int, projectName string, taskName string, status int, workingTime int) *Task {
	return &Task {
		id, projectName, taskName, status, workingTime,
	}
}

func (t *Task) NameTask() string {
	return fmt.Sprintf("%s: %s", t.ProjectName, t.TaskName)
}


func NewTaskFromDataItem(di binding.DataItem) *Task {
	v, _ := di.(binding.Untyped).Get()
	return v.(*Task)
}


type TaskDatas struct {
	binding.UntypedList
}

func NewTaskData (tasks []Task) TaskDatas {
	t := TaskDatas {
		binding.NewUntypedList(),
	}

	for _, task := range tasks {
		tasktemp := task
		t.Add(&tasktemp)
	}

	return t
}

func (t *TaskDatas) Add(task *Task) {
	
	t.Prepend(task)
}




var Tasks = []Task{
	*NewTask(1, "Project A", "Task 1", 1, 5),
	*NewTask(2, "Project B", "Task 1", 1, 6),
	// {"Project A", "Task 1", 0, 5},
	// {"Project A", "Task 2", 1, 8},
	// {"Project B", "Task 1", 0, 6},
	// {"Project B", "Task 2", 2, 7},
	// {"Project C", "Task 1", 1, 4},
	// {"Project C", "Task 2", 0, 9},
	// {"Project D", "Task 1", 2, 3},
	// {"Project D", "Task 2", 1, 6},
	// {"Project E", "Task 1", 0, 7},
	// {"Project E", "Task 2", 2, 5},
}

