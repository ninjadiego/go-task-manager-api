<div align="center">

# 🚀 Go Task Manager API

### A production-ready REST API built with Go, featuring JWT authentication, clean architecture, and modern DevOps practices.

[![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/build-passing-success?style=for-the-badge)](https://github.com/ninjadiego/go-task-manager-api)

</div>

---

## 📋 Table of Contents

- [Overview](#-overview)
- - [Features](#-features)
  - - [Tech Stack](#-tech-stack)
    - - [Architecture](#-architecture)
      - - [Getting Started](#-getting-started)
        - - [API Endpoints](#-api-endpoints)
          - - [Authentication](#-authentication)
            - - [Testing](#-testing)
              - - [Docker](#-docker)
                - - [Contributing](#-contributing)
                  - - [License](#-license)
                   
                    - ---

                    ## 🎯 Overview

                    **Go Task Manager API** is a robust, scalable REST API designed for managing tasks, projects, and team collaboration. Built following clean architecture principles, it provides a solid foundation for production-grade applications.

                    This project demonstrates best practices in Go development including dependency injection, repository pattern, middleware composition, structured logging, and comprehensive testing.

                    ## ✨ Features

                    - 🔐 **JWT Authentication** — Secure token-based authentication with refresh tokens
                    - - 👥 **User Management** — Registration, login, profile management
                      - - 📝 **Task CRUD** — Create, read, update, delete tasks with rich metadata
                        - - 🏷️ **Tags & Categories** — Organize tasks with custom labels
                          - - 🔍 **Advanced Filtering** — Search, sort, paginate, and filter results
                            - - 📊 **Analytics Dashboard** — Track productivity metrics
                              - - 🔔 **Notifications** — Email and webhook notifications
                                - - 📚 **Swagger Documentation** — Interactive API docs at `/swagger/index.html`
                                  - - 🐳 **Docker Ready** — One-command deployment with Docker Compose
                                    - - ⚡ **High Performance** — Built with Fiber/Gin for ultra-fast HTTP handling
                                      - - 🧪 **Tested** — Unit and integration tests with >80% coverage
                                        - - 📈 **Observability** — Structured logging, metrics, and tracing
                                         
                                          - ## 🛠️ Tech Stack
                                         
                                          - | Category | Technology |
                                          - |----------|-----------|
                                          - | **Language** | Go 1.22+ |
                                          - | **Framework** | Gin Web Framework |
                                          - | **Database** | PostgreSQL 16 |
                                          - | **ORM** | GORM |
                                          - | **Authentication** | JWT (golang-jwt/jwt) |
                                          - | **Validation** | go-playground/validator |
                                          - | **Documentation** | Swagger / OpenAPI 3.0 |
                                          - | **Testing** | Testify, GoMock |
                                          - | **Containerization** | Docker, Docker Compose |
                                          - | **CI/CD** | GitHub Actions |
                                          - | **Logging** | Zap (Uber) |
                                          - | **Configuration** | Viper |
                                         
                                          - ## 🏛️ Architecture
                                         
                                          - The project follows **Clean Architecture** principles with clear separation of concerns:
                                         
                                          - ```
                                            ┌─────────────────────────────────────────┐
                                            │         HTTP Handlers (Controllers)     │
                                            ├─────────────────────────────────────────┤
                                            │         Services (Business Logic)       │
                                            ├─────────────────────────────────────────┤
                                            │         Repositories (Data Access)      │
                                            ├─────────────────────────────────────────┤
                                            │         Database (PostgreSQL)           │
                                            └─────────────────────────────────────────┘
                                            ```

                                            ## 🚀 Getting Started

                                            ### Prerequisites

                                            - Go 1.22 or higher
                                            - - PostgreSQL 16+
                                              - - Docker & Docker Compose (optional)
                                                - - Make (optional)
                                                 
                                                  - ### Installation
                                                 
                                                  - ```bash
                                                    # Clone the repository
                                                    git clone https://github.com/ninjadiego/go-task-manager-api.git
                                                    cd go-task-manager-api

                                                    # Install dependencies
                                                    go mod download

                                                    # Copy environment variables
                                                    cp .env.example .env

                                                    # Run migrations
                                                    make migrate-up

                                                    # Start the server
                                                    make run
                                                    ```

                                                    The API will be available at `http://localhost:8080`

                                                    ## 📡 API Endpoints

                                                    ### Authentication
                                                    - `POST /api/v1/auth/register` — Register a new user
                                                    - - `POST /api/v1/auth/login` — Login and receive JWT
                                                      - - `POST /api/v1/auth/refresh` — Refresh access token
                                                        - - `POST /api/v1/auth/logout` — Invalidate token
                                                         
                                                          - ### Tasks
                                                          - - `GET /api/v1/tasks` — List all tasks (with pagination)
                                                            - - `POST /api/v1/tasks` — Create a new task
                                                              - - `GET /api/v1/tasks/:id` — Get task by ID
                                                                - - `PUT /api/v1/tasks/:id` — Update a task
                                                                  - - `DELETE /api/v1/tasks/:id` — Delete a task
                                                                   
                                                                    - ### Users
                                                                    - - `GET /api/v1/users/me` — Get current user profile
                                                                      - - `PUT /api/v1/users/me` — Update profile
                                                                        - - `DELETE /api/v1/users/me` — Delete account
                                                                         
                                                                          - ## 🔒 Authentication
                                                                         
                                                                          - This API uses JWT (JSON Web Tokens) for authentication. Include the token in the `Authorization` header:
                                                                         
                                                                          - ```
                                                                            Authorization: Bearer <your-jwt-token>
                                                                            ```

                                                                            ## 🧪 Testing

                                                                            ```bash
                                                                            # Run all tests
                                                                            make test

                                                                            # Run with coverage
                                                                            make test-coverage

                                                                            # Run integration tests
                                                                            make test-integration
                                                                            ```

                                                                            ## 🐳 Docker

                                                                            ```bash
                                                                            # Build and run with Docker Compose
                                                                            docker-compose up --build

                                                                            # Run in detached mode
                                                                            docker-compose up -d
                                                                            ```

                                                                            ## 🤝 Contributing

                                                                            Contributions are welcome! Please feel free to submit a Pull Request.

                                                                            ## 📝 License

                                                                            This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

                                                                            ---

                                                                            <div align="center">

                                                                            **Built with ❤️ by [ninjadiego](https://github.com/ninjadiego)**

                                                                            ⭐ Star this repo if you find it useful!

                                                                            </div>
                                                                            
