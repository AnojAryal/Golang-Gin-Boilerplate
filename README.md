## Go JWT Authentication Boilerplate

# This repository is a starter template for creating a Go API project with JWT authentication, CRUD operations, and a clean modular structure. It uses the following technologies:

1) Gin: Web framework
2) GORM: Object-Relational Mapper (ORM)
3) Postgres (or any other database supported by GORM)
4) JWT: JSON Web Token for authentication
5) godotenv: To load environment variables
6) CompileDaemon: For hot reloading during development(or any other as your preference)


## Features

1) Modular structure for scalability and readability
2) JWT-based authentication
3) CRUD operations for managing resources
4) Middleware for authentication and other logic
5) Environment variable management
6) Database connection and migration setup

## Project Structure


- ├── main.go                   # Entry point for the application
- ├── controllers/              # Handlers for routes (CRUD logic)
- │   ├── auth.go               # Handles authentication (login/register)
- │   ├── <other_entities>.go   # CRUD operations for other entities
- ├── initializers/             # Setup and configuration files
- │   ├── db.go                 # Database connection setup
- │   ├── env.go                # Load environment variables
- │   ├── migrate.go            # Sync database schema
- ├── models/                   # Database model definitions
- │   ├── user.go               # User model
- │   ├── <other_models>.go     # Other models
- ├── middleware/               # Middleware logic
- │   ├── auth.go               # JWT authentication middleware
- ├── routes/                   # API route definitions
- │   ├── routes.go             # Registers routes
- ├── .env                      # Environment variables
- ├── go.mod                    # Module definition
- ├── go.sum                    # Dependency lock file
- └── README.md                 # Project documentation
