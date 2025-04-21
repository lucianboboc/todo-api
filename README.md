# Todo API

A RESTful API for managing todo items with user authentication, built with Go.

## Overview

This Todo API provides a simple and efficient way to manage todo items. It includes user authentication, todo creation, retrieval, and management features. The API is built using Go and follows clean architecture principles.

## Features

- User registration and authentication
- JWT-based authentication
- Todo item management (create, retrieve)
- PostgreSQL database for data persistence
- Database migrations using Goose

## Prerequisites

- Go 1.22 or higher
- Docker and Docker Compose
- PostgreSQL (or use the provided Docker setup)

## Setup and Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/todo-api.git
   cd todo-api
   ```

2. Create a `.env` file in the root directory with the following content:
   ```
   PORT=8080
   DB_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
   JWT_SECRET=your_jwt_secret_key
   ```

3. Start the PostgreSQL database using Docker Compose:
   ```
   docker-compose up -d
   ```

4. Run database migrations:
   ```
   make migration-up
   ```

5. Start the API server:
   ```
   make start
   ```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Authentication

- **POST /api/v1/auth/login** - Authenticate a user
  ```json
  {
    "email": "user@example.com",
    "password": "password123"
  }
  ```
  Response:
  ```json
  {
    "token": "jwt_token_here"
  }
  ```

### Users

- **POST /api/v1/users** - Create a new user (register)
  ```json
  {
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "password": "password123"
  }
  ```

- **GET /api/v1/users/{id}** - Get a user by ID (requires authentication)

- **GET /api/v1/users** - Get all users (requires authentication)

- **GET /api/v1/users?email=user@example.com** - Get a user by email (requires authentication)

### Todos

- **POST /api/v1/todos** - Create a new todo (requires authentication)
  ```json
  {
    "text": "Buy groceries",
    "user_id": 1,
    "completed": false
  }
  ```

- **GET /api/v1/todos** - Get all todos (requires authentication)

- **GET /api/v1/todos/{id}** - Get a todo by ID (requires authentication)

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email citext UNIQUE NOT NULL,
    password_hash text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
```

### Todos Table

```sql
CREATE TABLE todos (
    id bigserial PRIMARY KEY,
    text text NOT NULL,
    completed bool NOT NULL DEFAULT false,
    user_id int,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## Available Commands

The project includes a Makefile with the following commands:

- `make start` - Start the API server
- `make new-migration name=<migration_name>` - Create a new migration file
- `make migration-up` - Apply all pending migrations
- `make migration-down` - Revert the last applied migration
- `make migration-status` - Check the status of migrations

## Authentication

The API uses JWT (JSON Web Token) for authentication. To access protected endpoints, include the JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Error Handling

The API returns appropriate HTTP status codes and error messages in JSON format using the `apierrors.APIError` type:

```json
{
  "error": {
    "domain": "error",
    "status_code": 400,
    "message": "Invalid user credentials",
    "key": "InvalidCredentials"
  }
}
```

## Development

The project follows a clean architecture pattern with the following structure:

- `cmd/api` - Application entry point
- `config` - Configuration loading
- `internal/domain` - Business logic and domain models
- `internal/infrastructure` - External services and tools
- `internal/transport` - HTTP handlers and middleware
- `migrations` - Database migration files

## 🤝 Contributing

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## License

[MIT License](LICENSE)
