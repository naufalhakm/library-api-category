# Category Microservice

## Overview
The **Category** Microservice is part of the **Library Management System**.

## Microservices

- **UserService/AuthService**: Handles user authentication and authorization.
- **BookService**: Manages books and stock.
- **CategoryService**: Manages book categories.
- **AuthorService**: Manages authors.

This microservice is built using:
- **Golang** for the backend.
- **PostgreSQL** for the database.
- **Docker** for containerization and deployment.
- **Docker Hub** for storing Docker images.

---

## **Technologies Used**
- **Programming Language**: Golang.
- **Database**: PostgreSQL.
- **Communication**: gRPC.
- **Middleware**: JWT for authentication.
- **Containerization**: Docker & Docker Compose.

---

## **API Documentation**
### REST API Endpoints
| HTTP Method | Endpoint                           | Description                          |
|-------------|------------------------------------|--------------------------------------|
| `GET`       | `/api/v1/categories`               | Get all categories                   |
| `POST`      | `/api/v1/categories`               | Create a new categories              |
| `GET`       | `/api/v1/categories/:id`           | Get details of a specific categories |
| `PUT`       | `/api/v1/categories/:id`           | Update a specific categories         |
| `DELETE`    | `/api/v1/categories/:id`           | Delete a specific categories         |
| `POST`      | `/api/v1/categories/books`         | Add book to categories               |
| `GET`       | `/api/v1/categories/books/:id`     | Get list categories of book          |

---

## Installation

### Prerequisites
- Install [Go](https://go.dev/doc/install)
- Install [PostgreSQL](https://www.postgresql.org/download/)
- Install [Docker](https://docs.docker.com/get-docker/)
- Install [gRPC](https://grpc.io/docs/languages/go/quickstart/)

### Running Without Docker

1. Clone the repository:
   ```sh
   git clone https://github.com/naufalhakm/library-api-category.git
   cd library-api-category
   ```
2. Setup environment variables (.env file):
   ```sh
   DB_HOST=localhost
   DB_PORT=5432
   DB_USERNAME=user
   DB_PASSWORD=password
   DB_DATABASE=library
   ```
3. Run PostgreSQL locally.
4. Start category microservice:
   ```sh
   go run cmd/server/main.go
   ```

### Running With Docker

1. Build and run services:
   ```sh
   docker-compose up -d
   ```

### Live Server

The microservice is running at:
http://35.240.139.186:8084/