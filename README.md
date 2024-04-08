# pitagoras-api

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Additional Commands](#additional-commands)
- [Usage](#usage)

## Introduction

This is a simple Go application that can be run on your local machine. It is a simple CRUD application that uses MongoDB as the database. It is built using the [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) and [Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) principles.

## Prerequisites

- [Go](https://golang.org/) (version 1.22.2 or later)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/) (optional)
- [Visual Studio Code](https://code.visualstudio.com/)

## Getting Started

- Clone the repository to your local machine: `git clone github.com/xbizzybone/pitagoras-api`
- Navigate to the project directory: `cd pitagoras-api`
- Install the dependencies: `go mod download`
- Build the application: `go build`
- Run the application: `.\pitagoras-api`

## Additional Commands

- To run the application in development mode with live reloading: `go run .`

## Usage

## API Endpoints

This application provides the following API endpoints:

- `POST /auth/register`: Register a new user
  - Request body: `UserCreateRequest`
  - Response body: `UserCreateResponse`

- `POST /auth/login`: Login a user
  - Request body: `UserRequest`
  - Response body: `UserLoginResponse`

- `GET /auth/user/:id`: Get a user by ID
  - Response body: `UserResponse`

- `GET /auth/user/email/:email`: Get a user by email
  - Response body: `UserResponse`

- `PUT /auth/user/activate/:id`: Activate a user
  - Response body: `UserResponse`

- `DELETE /auth/user/deactivate/:id`: Deactivate a user
  - Response body: `UserResponse`

- `PUT /auth/user/:id`: Update a user
  - Request body: `UserUpdateRequest`
  - Response body: `UserResponse`

### Request and Response Models

#### UserCreateRequest

```go
type UserCreateRequest struct {
    Name     string `bson:"name,omitempty" json:"name" validate:"required"`
    Email    string `bson:"email,omitempty" json:"email" validate:"required,email"`
    Password string `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
}
```

#### UserUpdateRequest

```go
type UserUpdateRequest struct {
    Name     string `bson:"name,omitempty" json:"name"`
    Email    string `bson:"email,omitempty" json:"email"`
    Password string `bson:"password,omitempty" json:"password,omitempty"`
}
```

#### UserRequest

```go
type UserRequest struct {
    Email     string             `bson:"email,omitempty" json:"email" validate:"required,email"`
    Password  string             `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
    LastLogin primitive.DateTime `bson:"last_login,omitempty" json:"last_login"`
    CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at" default:"now()"`
    UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
    DeleteAt  primitive.DateTime `bson:"delete_at,omitempty" json:"delete_at"`
}
```

#### UserResponse

```go
type UserResponse struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string             `bson:"name,omitempty" json:"name"`
    Email     string             `bson:"email,omitempty" json:"email"`
    Password  string             `bson:"password,omitempty" json:"password,omitempty"`
    IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
    LastLogin primitive.DateTime `bson:"last_login,omitempty" json:"last_login"`
    CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at" default:"now()"`
    UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
    DeleteAt  primitive.DateTime `bson:"delete_at,omitempty" json:"delete_at"`
}
```

#### UserLoginResponse

```go
type UserLoginResponse struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string             `bson:"name,omitempty" json:"name"`
    Email     string             `bson:"email,omitempty" json:"email"`
    Password  string             `bson:"password,omitempty" json:"password,omitempty"`
    LastLogin primitive.DateTime `bson:"last_login,omitempty" json:"last_login"`
    CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at" default:"now()"`
    UpdatedAt primitive.DateTime `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
}
```

#### UserCreateResponse

```go
type UserCreateResponse struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name  string             `bson:"name,omitempty" json:"name"`
    Email string             `bson:"email,omitempty" json:"email"`
}
```
