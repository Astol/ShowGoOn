# Web Application Project

This project is a web application that consists of a Go backend. The backend follows the repository pattern for data access and business logic.

## Project Structure

```
web-app-project
├── backend
│   ├── cmd
│   │   └── main.go
│   ├── internal
│   │   ├── repository
│   │   │   ├── slide_repository.go
│   │   │   └── settings_repository.go
│   │   ├── service
│   │   │   ├── slide_service.go
│   │   │   └── settings_service.go
│   │   └── handler
│   │       ├── slide_handler.go
│   │       └── settings_handler.go
│   ├── go.mod
│   └── go.sum
└── README.md
```

## Backend

The backend is built using Go and follows the repository pattern. It consists of:

- **cmd/main.go**: Entry point of the application, initializes the server and sets up routing.
- **internal/repository**: Contains the repository interfaces and methods for data access.
  - `slide_repository.go`: Manages database operations for slides.
  - `settings_repository.go`: Manages database operations for settings.
- **internal/service**: Implements business logic and interacts with the repository.
  - `slide_service.go`: Contains business logic for slides.
  - `settings_service.go`: Contains business logic for settings.
- **internal/handler**: Defines HTTP handlers for incoming requests.
  - `slide_handler.go`: Handles HTTP requests for slides.
  - `settings_handler.go`: Handles HTTP requests for settings.

## Setup Instructions

1. Clone the repository:
   ```
   git clone https://github.com/Astol/ShowGoOn
   cd ShowGoOn
   ```

2. Set up the backend:
   - Navigate to the `backend` directory.
   - Run `go mod tidy` to install dependencies.
   - Start the server with `go run cmd/main.go`.

## Usage

Access the application in your browser at `http://localhost:8080` for the backend (default port).

## License

This project is licensed under the MIT License.
