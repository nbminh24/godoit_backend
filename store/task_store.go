package store

import (
	"fmt"
	"godoit_backend/models"
	"sync"
)

// TaskStore handles the data logic for tasks.
// It uses a mutex to handle concurrent requests safely.
// Using a map for easier lookups and deletions.
type TaskStore struct {
	mu     sync.Mutex
	tasks  map[int]models.Task
	nextID int
}

// NewTaskStore creates and initializes a new TaskStore with sample data.
func NewTaskStore() *TaskStore {
	ts := &TaskStore{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}

	// Add initial data
	_, _ = ts.Create("Học Golang", "Tìm hiểu về Gin framework")
	_, _ = ts.Create("Làm bài test", "Hoàn thành phần backend")
	task3, _ := ts.Create("Triển khai dự án", "Deploy lên Render")

	// Manually update the last task to be completed for sample data
	if task, ok := ts.tasks[task3.ID]; ok {
		task.Completed = true
		ts.tasks[task.ID] = task
	}

	return ts
}

// GetAll returns all tasks from the store.
func (ts *TaskStore) GetAll() []models.Task {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	allTasks := make([]models.Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

// Create adds a new task to the store.
func (ts *TaskStore) Create(title, description string) (models.Task, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	newTask := models.Task{
		ID:          ts.nextID,
		Title:       title,
		Description: description,
		Completed:   false,
	}

	ts.tasks[newTask.ID] = newTask
	ts.nextID++
	return newTask, nil
}

// FindByID finds a task by its ID.
func (ts *TaskStore) FindByID(id int) (models.Task, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task, ok := ts.tasks[id]
	if !ok {
		return models.Task{}, fmt.Errorf("task with ID %d not found", id)
	}
	return task, nil
}

// Update updates a task's completed status.
func (ts *TaskStore) Update(id int, completed bool) (models.Task, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	task, ok := ts.tasks[id]
	if !ok {
		return models.Task{}, fmt.Errorf("task with ID %d not found", id)
	}

	task.Completed = completed
	ts.tasks[id] = task
	return task, nil
}

// Delete removes a task from the store by its ID.
func (ts *TaskStore) Delete(id int) error {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	_, ok := ts.tasks[id]
	if !ok {
		return fmt.Errorf("task with ID %d not found", id)
	}

	delete(ts.tasks, id)
	return nil
}
