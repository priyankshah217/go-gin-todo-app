# Go Gin Todo App

A sample RESTful Todo application built with Go and Gin framework, featuring in-memory database, Docker support, and complete CRUD operations.

## Features

- RESTful API endpoints for Todo management
- In-memory data storage
- CORS support
- Health check endpoint
- Docker containerization
- JSON API responses

## Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerization)

## Installation

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/priyankshah217/go-gin-todo-app.git
   cd go-gin-todo-app
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

### Running with Docker

1. Build the Docker image:
   ```bash
   docker build -t go-gin-todo-app .
   ```

2. Run the container:
   ```bash
   docker run -p 8080:8080 go-gin-todo-app
   ```

The server will be available at `http://localhost:8080`

## API Endpoints

### Health Check
- **GET** `/health` - Returns service health status

### Todo Operations

All todo endpoints are prefixed with `/api/v1`

#### Create a Todo
- **POST** `/api/v1/todos`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "title": "Sample Todo",
    "description": "This is a sample todo item",
    "completed": false
  }
  ```
- **Response**: `201 Created`
  ```json
  {
    "id": 1,
    "title": "Sample Todo",
    "description": "This is a sample todo item",
    "completed": false,
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:00:00Z"
  }
  ```

#### Get All Todos
- **GET** `/api/v1/todos`
- **Response**: `200 OK`
  ```json
  {
    "todos": [
      {
        "id": 1,
        "title": "Sample Todo",
        "description": "This is a sample todo item",
        "completed": false,
        "created_at": "2023-01-01T12:00:00Z",
        "updated_at": "2023-01-01T12:00:00Z"
      }
    ]
  }
  ```

#### Get a Specific Todo
- **GET** `/api/v1/todos/{id}`
- **Response**: `200 OK`
  ```json
  {
    "id": 1,
    "title": "Sample Todo",
    "description": "This is a sample todo item",
    "completed": false,
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:00:00Z"
  }
  ```
- **Response**: `404 Not Found` (if todo doesn't exist)
  ```json
  {
    "error": "Todo not found"
  }
  ```

#### Update a Todo
- **PUT** `/api/v1/todos/{id}`
- **Content-Type**: `application/json`
- **Request Body**:
  ```json
  {
    "title": "Updated Todo",
    "description": "This is an updated todo item",
    "completed": true
  }
  ```
- **Response**: `200 OK`
  ```json
  {
    "id": 1,
    "title": "Updated Todo",
    "description": "This is an updated todo item",
    "completed": true,
    "created_at": "2023-01-01T12:00:00Z",
    "updated_at": "2023-01-01T12:05:00Z"
  }
  ```
- **Response**: `404 Not Found` (if todo doesn't exist)

#### Delete a Todo
- **DELETE** `/api/v1/todos/{id}`
- **Response**: `200 OK`
  ```json
  {
    "message": "Todo deleted successfully"
  }
  ```
- **Response**: `404 Not Found` (if todo doesn't exist)
  ```json
  {
    "error": "Todo not found"
  }
  ```

## Example Usage

### Using curl

1. **Create a todo**:
   ```bash
   curl -X POST http://localhost:8080/api/v1/todos \
     -H "Content-Type: application/json" \
     -d '{"title":"Learn Go","description":"Complete Go tutorial","completed":false}'
   ```

2. **Get all todos**:
   ```bash
   curl http://localhost:8080/api/v1/todos
   ```

3. **Get a specific todo**:
   ```bash
   curl http://localhost:8080/api/v1/todos/1
   ```

4. **Update a todo**:
   ```bash
   curl -X PUT http://localhost:8080/api/v1/todos/1 \
     -H "Content-Type: application/json" \
     -d '{"title":"Learn Go - Updated","description":"Complete Go tutorial and build an app","completed":true}'
   ```

5. **Delete a todo**:
   ```bash
   curl -X DELETE http://localhost:8080/api/v1/todos/1
   ```

6. **Health check**:
   ```bash
   curl http://localhost:8080/health
   ```

## Project Structure

```
go-gin-todo-app/
├── main.go           # Main application file with API endpoints
├── go.mod           # Go module dependencies
├── Dockerfile       # Docker configuration
└── README.md        # This file
```

## Technology Stack

- **Go 1.21**: Programming language
- **Gin**: Web framework for Go
- **In-memory storage**: Simple data persistence
- **Docker**: Containerization

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).
