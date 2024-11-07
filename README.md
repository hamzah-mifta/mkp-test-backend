# Golang CRUD API

A simple RESTful API for managing movie schedules in theaters. Built using Go and Echo framework, with PostgreSQL as the database.

This repository is a part of recruitment process for Backend Developer position in PT. Mitra Kasih Perkasa. The answers of written test (flowchart, explanation, database design and postman colection) located in `docs` directory

## Technologies Used

- **Go** (Golang)
- **Echo**: Web framework for building RESTful APIs.
- **PostgreSQL**: Database for storing data.
- **golang-migrate**: Tool for database migrations.

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (version 1.18 or later)
- [PostgreSQL](https://www.postgresql.org)
- [golang-migrate](https://github.com/golang-migrate/migrate)

### Installation

1. Clone the repository

   ```bash
   git clone git@github.com:hamzah-mifta/mkp-test-backend.git
   ```

2. Install dependencies

   ```bash
   go mod tidy
   ```

3. Run database migration
   ```bash
   migrate -path migrations -database 'postgres://username:password@localhost:5432/dbname?sslmode=disable' up
   ```

### Running the Application

```bash
go run cmd/api/main.go
```

## Lisence

This project is licensed under the MIT License. See `LICENSE` for more information.

## Author

[Hamzah Miftakhuddin](https://github.com/hamzah-mifta)
