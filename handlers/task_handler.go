package handlers

import (
	"godoit_backend/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TaskHandler connects the API routes to the data store.
type TaskHandler struct {
	store *store.TaskStore
}

// NewTaskHandler creates a new handler with a given task store.
func NewTaskHandler(ts *store.TaskStore) *TaskHandler {
	return &TaskHandler{store: ts}
}

// GetTasks responds with the list of all tasks as JSON.
// GET /tasks
func (th *TaskHandler) GetTasks(c *gin.Context) {
	tasks := th.store.GetAll()
	c.IndentedJSON(http.StatusOK, tasks)
}

// CreateTaskRequest defines the shape of the request body for creating a task.
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// CreateTask adds a new task from JSON received in the request body.
// POST /tasks
func (th *TaskHandler) CreateTask(c *gin.Context) {
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := th.store.Create(req.Title, req.Description)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	// Respond with the ID of the newly created task
	c.IndentedJSON(http.StatusCreated, gin.H{"id": newTask.ID})
}

// UpdateTaskRequest defines the shape of the request body for updating a task's status.
type UpdateTaskRequest struct {
	Completed bool `json:"completed"`
}

// UpdateTask updates a task's status.
// PUT /tasks/:id
func (th *TaskHandler) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedTask, err := th.store.Update(id, req.Completed)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedTask)
}

// DeleteTask deletes a task by ID.
// DELETE /tasks/:id
func (th *TaskHandler) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := th.store.Delete(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

