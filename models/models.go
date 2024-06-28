package models


type Task struct {
	ProjectName string
	TaskName string
	Status int
	WorkingTime int
}


var Tasks = []Task{
	{"Project A", "Task 1", 0, 5},
	// {"Project A", "Task 2", 1, 8},
	{"Project B", "Task 1", 0, 6},
	// {"Project B", "Task 2", 2, 7},
	// {"Project C", "Task 1", 1, 4},
	// {"Project C", "Task 2", 0, 9},
	// {"Project D", "Task 1", 2, 3},
	// {"Project D", "Task 2", 1, 6},
	// {"Project E", "Task 1", 0, 7},
	// {"Project E", "Task 2", 2, 5},
}