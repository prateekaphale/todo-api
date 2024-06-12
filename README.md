

This API allows you to manage TODO items. You can create, read, update, and delete TODO items, as well as list them with pagination.

## Getting Started

To run the TODO List API, follow these steps:

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/todo-list-api.git
   cd todo-list-api
2.Install dependencies:


go mod tidy

3.Run the application:


go run main.go

Access the API at http://localhost:8080

API Endpoints
Create TODO: POST /todos
Request Body: JSON with title and description fields
Get TODO: GET /todos/{id}
Update TODO: PUT /todos/{id}
Request Body: JSON with title and description fields
Delete TODO: DELETE /todos/{id}
List TODOs: GET /todos
Query Parameters: page and limit for pagination
Design Decisions
Gorilla Mux: Chosen for its simplicity and powerful routing capabilities.
ScyllaDB: Used for its scalability and performance in handling large amounts of data.
Pagination: Implemented to improve API performance and usability for large datasets.
Error Handling: 400 Bad Request for invalid requests, 404 Not Found for non-existing resources, and 500 Internal Server Error for server-side issues.