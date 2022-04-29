package main

import "time"

// Some basic data types that our application uses

// Task represents a task for the to-do list
type Task struct {
	ID          int
	Description string
	Due         time.Time
	Completed   bool
}

// Generate a list of basic tasks for demo purposes
func generateTasks() []Task {
	return []Task{
		{
			ID:          1,
			Description: "Get out of bed",
			Due:         time.Now().Add(-time.Hour * 2),
			Completed:   true,
		},
		{
			ID:          2,
			Description: "Decide on lunch",
			Due:         time.Now().Add(time.Hour),
			Completed:   false,
		},
		{
			ID:          3,
			Description: "Eat lunch",
			Due:         time.Now().Add(time.Hour * 2),
			Completed:   false,
		},
		{
			ID:          4,
			Description: "Sleep",
			Due:         time.Now().Add(time.Hour * 18),
			Completed:   false,
		},
	}
}
