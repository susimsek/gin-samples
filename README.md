# Gin Samples

![Build Status](https://github.com/susimsek/gin-samples/actions/workflows/deploy.yml/badge.svg)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=gin-samples&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=gin-samples)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=gin-samples&metric=coverage)](https://sonarcloud.io/summary/new_code?id=gin-samples)
![Top Language](https://img.shields.io/github/languages/top/susimsek/gin-samples)

Welcome to **Gin Samples** – a lightweight and fast HTTP server built using the Gin framework.

Explore high performance, easy-to-use routing, and flexible middleware with Gin. 🚀

## 🚀 Quick Links

- 📖 [Features](#-features)
- 🧑‍💻 [Development Setup](#-development-setup)
- 🔄 [Live Reload](#-live-reload)
- 🧪 [Testing](#-testing)
- 🏗️ [Build](#️-build)
- 🕵️‍♂️ [Code Analysis](#️-code-analysis)
- 🛡️ [Code Quality](#️-code-quality)
- 📜 [API Documentation](#-api-documentation)
- 🐳 [Docker](#-docker)
- 🛠️ [Used Technologies](#️-used-technologies)

## 📖 Features

- 🚀 **High Performance**: Built for speed and efficiency.
- 🌐 **Simple Routing**: Define routes with minimal code.
- 🔌 **Middleware Support**: Easily add middleware to your application.
- 🤪 **Extensible**: Add your own routes and features.

## 🧑‍💻 Development Setup

To clone and run this application locally:

```bash
# Clone the repository
git clone https://github.com/example/gin-samples.git

# Navigate to the project directory
cd gin-samples

# Install dependencies (if required)
go mod tidy

# Start the application
go run cmd/app/main.go
```

## 🔄 Live Reload

`Air` is a live reload tool for Go applications that automatically rebuilds and restarts the application whenever you make changes to the codebase. It's ideal for speeding up development workflows.

### Using Air

1. Ensure `Air` is installed in your system:
   ```bash
   go install github.com/air-verse/air@latest
   ```

2. 	(Optional) Install Delve for debugging:
   ```bash
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

3.  Run the application with Air:
   ```bash
   air
   ```

When you make changes to your code, Air will automatically rebuild and restart the application. If you’ve set up Delve, you can debug on localhost:2345.

## 🧪 Testing

To test the application:

```bash
curl -X GET http://localhost:8080/api/hello
```

Expected response:

```json
{
  "message": "Hello, World!"
}
```

### Running Unit Tests

Run the following command to execute unit tests:

```bash
go test ./... -v -cover
```

## 🏗️ Build

To build the application for production:

```bash
go build -o gin-samples
```

## 🕵️ Code Analysis

`GolangCI-Lint` is a powerful and fast linters runner for Go. It helps maintain code quality by analyzing the source code for potential issues and providing suggestions for improvement.

### Setup GolangCI-Lint

1. Install GolangCI-Lint:
   ```bash
   brew install golangci/tap/golangci-lint
   ```
   Or, for a manual installation:
   ```bash
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
   ```

3. Run GolangCI-Lint:
   ```bash
   golangci-lint run
   ```

By integrating GolangCI-Lint into your Go projects, you ensure consistent and high-quality code, saving time and effort in debugging and maintenance.

## 🛡️ Code Quality

To assess code quality locally using SonarQube, execute:

```bash
sonar-scanner
```

## 📜 API Documentation

To view the API documentation, access the Swagger UI at:

http://localhost:8080/swagger-ui/index.html

## 🐳 Docker

To build and run the application using Docker:

### Build Docker Image

```bash
docker build -t gin-samples .
```

### Run Docker Container

```bash
docker run -d -p 8080:8080 gin-samples
```

The application will be available at `http://localhost:8080`.

## 🛠️ Used Technologies

![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go&logoColor=white)  
![Gin Framework](https://img.shields.io/badge/Gin_Framework-000000?logo=go&logoColor=white)  
![Gin Swagger](https://img.shields.io/badge/Gin_Swagger-85EA2D?logo=swagger&logoColor=white)  
![Go Playground Validator](https://img.shields.io/badge/Go_Playground_Validator-FDD835?logo=go&logoColor=white)  
![GORM](https://img.shields.io/badge/GORM-Go_ORM-7462FF?logo=go&logoColor=white)  
![SQLite](https://img.shields.io/badge/SQLite-003B57?logo=sqlite&logoColor=white)  
![Golang Migrate](https://img.shields.io/badge/Golang_Migrate-Database_Migrations-0E83CD?logo=go&logoColor=white)  
![SonarQube](https://img.shields.io/badge/SonarQube-4E9BCD?logo=sonarqube&logoColor=white)  
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)  
![Air](https://img.shields.io/badge/Air-Live_Reload-green?logo=go&logoColor=white)  
![GolangCI-Lint](https://img.shields.io/badge/GolangCI--Lint-Code_Analysis-orange?logo=go&logoColor=white)  
![Testify](https://img.shields.io/badge/Testify-Mocking_Framework-6E85B7?logo=go&logoColor=white)

---

This project is an excellent starting point for building web applications with the Gin framework!
