# Node.js Express Todo API with Docker

## @Siddharth9890

## DESCRIPTION OF PROJECT
```
This is a simple Todo API built with Node.js and Express, containerized with Docker.
```

## Prerequisites
Docker

## Steps to run

- Build the docker image for the TODO api

```bash
docker build -t todo-app .

```

- Run the container 
```bash
docker run -p 3000:3000 -d todo-app
```

## API Endpoints
# Create a Task

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "taskName": "Sample Task",
  "description": "This is a sample task."
}' http://localhost:3000/tasks

```
# Get All Tasks
```bash
curl http://localhost:3000/tasks

```

# Update a Task
```bash
curl -X PUT -H "Content-Type: application/json" -d '{
  "taskName": "Updated Task",
  "description": "This task has been updated."
}' http://localhost:3000/tasks/{taskId}

```

# Delete a Task
```bash
curl -X DELETE http://localhost:3000/tasks/{taskId}

```