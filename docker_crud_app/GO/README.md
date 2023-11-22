# Go Backend Application

This is a simple Go API implemented using the Gin web framework. It provides basic CRUD (Create, Read, Update, Delete).

## Table of Contents

- [Go Backend Application](#go-backend-application)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Getting Started](#getting-started)
  - [API Endpoints](#api-endpoints)
  - [Contributor @pushkarm029](#contributor-pushkarm029)

## Prerequisites

Before you begin, ensure you have the following dependencies installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/DiptoChakrabarty/learn_devops_with_projects
   ```

2. Change into the project directory:

   ```bash
   cd learn_devops_with_projects/docker_crud_app/go
   ```

3. Build and run the Go API:

   ```bash
   docker compose up
   ```

   This will start the API server on port 8080 by default.

## API Endpoints

The API provides the following endpoints:

- **GET /get**: Get the list of storage items.

- **POST /post**: Create a new storage item.
  - Request body should be in JSON format with a "text" field.
  - Example:
    ```json
    {
      "text": "New storage item"
    }
    ```

- **PUT /put**: Update an existing storage item.
  - Request body should be in JSON format with an "id" (the ID of the item to update) and a "text" field.
  - Example:
    ```json
    {
      "id": 1,
      "text": "Updated text"
    }
    ```

- **DELETE /delete**: Delete a storage item.
  - Request body should be in JSON format with an "id" (the ID of the item to delete).
  - Example:
    ```json
    {
      "id": 1
    }
    ```

Now you can access the API at [http://localhost:8080](http://localhost:8080).

## Contributor @pushkarm029