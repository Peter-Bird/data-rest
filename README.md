## WF-DBA

WF-DBA is a Go-based application designed to manage workflows and tasks. This application includes a series of scripts and components for managing workflows through database interactions, repository handling, and service management.

## Features

- **Workflow Management**: Easily submit, retrieve, list, and truncate workflows using the provided shell scripts.
- **Modular Design**: Organized with a clear package structure for handling database connections, repositories, services, request handling, and routing.
- **Configuration Management**: Centralized configuration setup for easier maintenance and scalability.

## Project Structure

- **main.go**: The main entry point of the application.
- **pkg/**: Contains core application packages:
  - `database.go`: Manages database connections and queries.
  - `repo.go`: Contains repository interfaces for interacting with data.
  - `services.go`: Defines services that encapsulate business logic.
  - `handlers.go`: Handles HTTP requests.
  - `routers.go`: Configures application routing.
  - `config.go`: Manages application configurations.
- **scripts/**: Utility scripts for common database and workflow tasks:
  - `submit.sh`: Submits a new workflow.
  - `get.sh`: Retrieves an existing workflow.
  - `list.sh`: Lists all workflows.
  - `truncate.sh`: Clears workflows from the database.

## Prerequisites

- Go 1.23.2 or higher.
- Bash (for running utility scripts).

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd wf-dba
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Usage

1. **Starting the Application**:
   ```bash
   go run main.go
   ```

2. **Using Scripts**:
   - Submit a workflow: `./scripts/submit.sh`
   - Get a workflow: `./scripts/get.sh`
   - List all workflows: `./scripts/list.sh`
   - Truncate workflows: `./scripts/truncate.sh`

3. **Using End Points**:

The `wf-dba` application provides a set of RESTful endpoints for managing workflows. Below is a list of available endpoints and their purposes.

### Base URL

By default, the application runs on `http://localhost:8083`. Update this in `config.go` if you need to change the port or host.

### Endpoints

#### 1. **Submit a Workflow**
   - **Endpoint**: `POST /api/workflows`
   - **Description**: Submits a new workflow to the system.
   - **Request Body**:
     ```json
     {
       "name": "Workflow Name",
       "description": "Detailed description of the workflow",
       "tasks": [
         {
           "task_name": "Task 1",
           "parameters": { "key": "value" }
         },
         {
           "task_name": "Task 2",
           "parameters": { "key": "value" }
         }
       ]
     }
     ```
   - **Response**:
     - `201 Created` with the workflow ID on successful submission.
     - `400 Bad Request` if the request is invalid.

#### 2. **Get Workflow by ID**
   - **Endpoint**: `GET /api/workflows/{id}`
   - **Description**: Retrieves the details of a specific workflow by its ID.
   - **Path Parameter**:
     - `id` (string): The unique identifier of the workflow.
   - **Response**:
     - `200 OK` with workflow details in JSON format.
     - `404 Not Found` if the workflow ID does not exist.

#### 3. **List All Workflows**
   - **Endpoint**: `GET /api/workflows`
   - **Description**: Returns a list of all workflows in the system.
   - **Response**:
     - `200 OK` with an array of workflows in JSON format.

#### 4. **Truncate Workflows**
   - **Endpoint**: `DELETE /api/workflows`
   - **Description**: Deletes all workflows from the database.
   - **Response**:
     - `200 OK` on successful deletion.
     - `500 Internal Server Error` if an issue occurs during deletion.

### Error Handling

- **400 Bad Request**: Invalid request payload or parameters.
- **404 Not Found**: Resource not found.
- **500 Internal Server Error**: Server encountered an error during request processing.

### Example Usage

Here’s an example of how to use the `curl` command to interact with the API endpoints.

1. **Submit a Workflow**:
   ```bash
   curl -X POST http://localhost:8080/api/workflows \
   -H "Content-Type: application/json" \
   -d '{
         "name": "Sample Workflow",
         "description": "A sample workflow example",
         "tasks": [
           {"task_name": "Task 1", "parameters": {"param1": "value1"}},
           {"task_name": "Task 2", "parameters": {"param2": "value2"}}
         ]
       }'


## Configuration

All configurations are managed in `config.go`. Update this file to change database connections or other environment settings.

## Contributing

Feel free to fork and submit pull requests. For significant changes, please open an issue to discuss.

## License

This project is licensed under the MIT License.
