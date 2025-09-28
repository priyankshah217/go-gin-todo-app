package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Todo represents a todo item
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// In-memory database
var (
	todos   []Todo
	nextID  int = 1
	todoMu  sync.RWMutex
)

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoMu.Lock()
	defer todoMu.Unlock()

	newTodo.ID = nextID
	nextID++
	newTodo.CreatedAt = time.Now()
	newTodo.UpdatedAt = time.Now()
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// GetTodos returns todos with pagination support
func GetTodos(c *gin.Context) {
	todoMu.RLock()
	defer todoMu.RUnlock()

	// Parse pagination parameters
	page := 1
	limit := 10

	if pageParam := c.Query("page"); pageParam != "" {
		if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
			page = p
		}
	}

	if limitParam := c.Query("limit"); limitParam != "" {
		if l, err := strconv.Atoi(limitParam); err == nil && l > 0 {
			if l <= 100 {
				limit = l
			} else {
				limit = 100 // Cap at maximum
			}
		}
	}

	// Calculate pagination
	totalCount := len(todos)
	totalPages := (totalCount + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	// Calculate offset
	offset := (page - 1) * limit
	if offset >= totalCount {
		// Page is beyond available data
		c.JSON(http.StatusOK, gin.H{
			"todos":        []Todo{},
			"total_count":  totalCount,
			"current_page": page,
			"total_pages":  totalPages,
			"per_page":     limit,
			"has_next":     false,
			"has_prev":     page > 1,
		})
		return
	}

	// Get the slice for this page
	end := offset + limit
	if end > totalCount {
		end = totalCount
	}

	paginatedTodos := todos[offset:end]

	c.JSON(http.StatusOK, gin.H{
		"todos":        paginatedTodos,
		"total_count":  totalCount,
		"current_page": page,
		"total_pages":  totalPages,
		"per_page":     limit,
		"has_next":     page < totalPages,
		"has_prev":     page > 1,
	})
}

// GetTodo returns a specific todo by ID
func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todoMu.RLock()
	defer todoMu.RUnlock()

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// UpdateTodo updates an existing todo
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todoMu.Lock()
	defer todoMu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			updatedTodo.ID = id
			updatedTodo.CreatedAt = todo.CreatedAt
			updatedTodo.UpdatedAt = time.Now()
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todoMu.Lock()
	defer todoMu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/todos", CreateTodo)
		v1.GET("/todos", GetTodos)
		v1.GET("/todos/:id", GetTodo)
		v1.PUT("/todos/:id", UpdateTodo)
		v1.DELETE("/todos/:id", DeleteTodo)
	}

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "go-gin-todo-app",
			"version": "1.0.0",
		})
	})

	// Start server on port 8080
	r.Run(":8080")
}
