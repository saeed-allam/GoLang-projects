# Bookstore Management APIs (Beginner Friendly)

This document outlines the API endpoints for managing a bookstore.

## Technologies Used

1.  Database - MySQL 
   ```bash
   go get -u github.com/jinzhu/gorm/dialects/mysql
2.  GORM (Go Object-Relational Mapping)
   ```bash
   go get -u github.com/jinzhu/gorm
3.  JSON (for request and response body serialization/deserialization)
4.  Project Structure (as detailed below)
5.  Gorilla Mux (HTTP request router and dispatcher)
   ```bash
   go get -u github.com/gorilla/mux

## API Endpoints

The following endpoints are available for interacting with the bookstore:

| HTTP Method | Endpoint          | Description        |
| :---------- | :---------------- | :----------------- |
| `GET`       | `/book/`          | Get all books      |
| `POST`      | `/book/`          | Create a new book  |
| `GET`       | `/book/{bookId}`  | Get a book by ID   |
| `PUT`       | `/book/{bookId}`  | Update a book      |
| `DELETE`    | `/book/{bookId}`  | Delete a book      |

**Note:** `{bookId}` is a path parameter representing the unique identifier of a book.

## Project Structure
CMD
└── main.go

PKG
├── config
│   └── app.go
├── controllers
│   └── book-controller.go
├── models
│   └── book.go
├── routes
│   └── bookstore-routes.go
└── utils
    └── utils.go

**Description of Directories and Files:**

* **`CMD`**: Contains the main application entry point.
    * `main.go`: The main Go file to run the application.
* **`PKG`**: Contains the application's packages (modules).
    * **`config`**: Configuration related files.
        * `app.go`: Likely contains application-wide configurations.
    * **`controllers`**: Handles the application logic for incoming requests.
        * `book-controller.go`: Contains the controller functions for book-related operations.
    * **`models`**: Defines the data structures used in the application.
        * `book.go`: Defines the structure of a book.
    * **`routes`**: Defines the API endpoints and their corresponding handlers.
        * `bookstore-routes.go`: Defines the routes for the bookstore API.
    * **`utils`**: Contains utility functions used across the application.
        * `utils.go`: Likely contains reusable utility functions.