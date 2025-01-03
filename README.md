# Gin Samples

Welcome to **Gin Samples** – a lightweight and fast HTTP server built using the Gin framework.

Explore high performance, easy-to-use routing, and flexible middleware with Gin. 🚀

## 🚀 Quick Links

- 📖 [Features](#-features)
- 🧑‍💻 [Development Setup](#-development-setup)
- 🧪 [Testing](#-testing)
- 🏗️ [Build](#️-build)
- 🛡️ [Code Quality](#️-code-quality)
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

## 🛡️ Code Quality

To assess code quality locally using SonarQube, execute:

```bash
sonar-scanner
```

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
![SonarQube](https://img.shields.io/badge/SonarQube-4E9BCD?logo=sonarqube&logoColor=white)  
![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)

---

This project is an excellent starting point for building web applications with the Gin framework!
