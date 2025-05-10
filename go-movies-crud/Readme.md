# Movies API

This is a simple Movies REST API built with Go using the Gorilla Mux router. The API runs on `localhost:8000` and interacts with a connected database to perform CRUD operations.

## Architecture Overview

The server uses Gorilla Mux for routing and connects to a database to manage movie data. Below are the available routes, their corresponding functions, endpoints, and HTTP methods.

## API Routes

| Route       | Function       | Endpoint      | HTTP Method |
|-------------|----------------|---------------|-------------|
| Get All     | `getMovies`    | `/movies`     | GET         |
| Get By ID   | `getMovie`     | `/movies/id`  | GET         |
| Create      | `createMovie`  | `/movies`     | POST        |
| Update      | `updateMovie`  | `/movies/id`  | PUT         |
| Delete      | `deleteMovie`  | `/movies/id`  | DELETE      |

## Tech Stack

- **Backend**: Go (Golang)
- **Router**: Gorilla Mux
- **Database**: Any standard DB supported by Go (e.g., PostgreSQL, MySQL)

## Running the Server

1. Make sure you have Go installed.
2. Initialize Go module (if not already):
   ```bash
   go mod init movies-api

3. Install Gorilla Mux:
   ```bash
   go get -u github.com/gorilla/mux

