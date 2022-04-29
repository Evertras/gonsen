package main

import "fmt"

// Repository represents a data source of some sort... in this case just a simple
// in-memory data store, but an actual app would have better data backing it.
type Repository struct {
	tasks []Task
}

func NewRepository() *Repository {
	return &Repository{
		tasks: generateTasks(),
	}
}

func (r *Repository) GetTasks() ([]Task, error) {
	return r.tasks, nil
}

func (r *Repository) GetTask(id int) (Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("task with id %d not found", id)
}

func (r *Repository) MarkTaskComplete(id int) error {
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks[i].Completed = true
			return nil
		}
	}

	return fmt.Errorf("task with id %d not found", id)
}
