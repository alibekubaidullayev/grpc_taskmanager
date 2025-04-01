
# Task Manager

A simple task management system built with **Go**, **Gorm**, **GRPC**, **PostgreSQL**, and **Docker**. This application allows users to perform basic CRUD operations (Create, Read, Update, Delete) on tasks.

## Technologies Used

- **Go**
- **Gorm**
- **GRPC**
- **PostgreSQL**
- **Docker**

## Task Model (Gorm)

The Task model contains the following fields:

- **ID**: Unique identifier for each task.
- **Created At**: Timestamp when the task was created.
- **Updated At**: Timestamp when the task was last updated.
- **Deleted At**: Timestamp when the task was deleted (soft delete).
- **Title**: Title of the task.
- **Description**: Detailed description of the task.

```go
type Task struct {
    ID          uint      `gorm:"primaryKey"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   *time.Time
    Title       string    `gorm:"size:64"`
    Description string    `gorm:"size:512"`
}
```
## Features
### 1. Each method returns only required information
e.g. Upon update only the updated fields and update_at will be returned

### 2. Data structure level validation
Each method tries to put all coming fields into struct using setters. If field doesn't apply then error will be returned and operation aborted

### 3. Application layers separation (if better paraphrase)
There is Task struct that is responible for working with db. It has methods for converting (and validating) data that comes from GRPC to gorm.Model.

### 4. Protobuf modification
If you want to change the GRPC messages then modify src/proto/task_manager.proto and in root directory do
```bash
make generate
```
To generate requried files


## GRPC Service Methods

### 1. **Create**
Creates a new task.

#### Request Object:
```protobuf
message CreateTaskRequest {
    string title = 1;
    optional string description = 2;
}
```

#### Response Object:
```protobuf
message GetTaskResponse {
    uint64 id = 1;
    string title = 2;
    string description = 3;
    google.protobuf.Timestamp created_at = 4;
    google.protobuf.Timestamp updated_at = 5;
    google.protobuf.Timestamp deleted_at = 6;
}
```

#### Screen Link:
[Link to Screen](#)

---

### 2. **Get**
Fetches a specific task by ID.

#### Request Object:
```protobuf
message IdRequest {
    uint64 id = 1;
}
```

#### Response Object:
```protobuf
message GetTaskResponse {
    uint64 id = 1;
    string title = 2;
    string description = 3;
    google.protobuf.Timestamp created_at = 4;
    google.protobuf.Timestamp updated_at = 5;
    google.protobuf.Timestamp deleted_at = 6;
}
```

#### Screen Link:
[Link to Screen](#)

---

### 3. **Update**
Updates an existing task.

#### Request Object:
```protobuf
message UpdateTaskRequest {
    uint64 id = 1;
    optional string title = 2;
    optional string description = 3;
}
```

#### Response Object:
```protobuf
message UpdateTaskResponse {
    uint64 id = 1;
    optional string title = 2;
    optional string description = 3;
    google.protobuf.Timestamp updated_at = 4;
}
```

#### Screen Link:
[Link to Screen](#)

---

### 4. **Delete**
Deletes a task by ID (soft delete).

#### Request Object:
```protobuf
message IdRequest {
    uint64 id = 1;
}
```

#### Response Object:
```protobuf
message GetTaskResponse {
    uint64 id = 1;
    string title = 2;
    string description = 3;
    google.protobuf.Timestamp created_at = 4;
    google.protobuf.Timestamp updated_at = 5;
    google.protobuf.Timestamp deleted_at = 6;
}
```

#### Screen Link:
[Link to Screen](#)

---

### 5. **List**
Fetches all tasks.

#### Request Object:
```
Empty
```

#### Response Object:
```protobuf
message ListTasksResponse {
    repeated GetTaskResponse tasks = 1;
}
```

#### Screen Link:
[Link to Screen](#)

---

## Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/alibekubaidullayev/task-manager.git
cd task-manager
```

### 3. Build and Run with Docker
To run the application in a Docker container, follow these steps:

```bash
docker-compose up --build
```
This will create the necessary tables in PostgreSQL using Gorm.

### 4. Application
Available by port 8081 on local device (modify docker-compose to change)


