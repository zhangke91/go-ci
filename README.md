# Simple Go Web Server

This is a simple web server written in Go that responds to GET requests.

## Running Locally

To run the server locally:

```bash
go run main.go
```

The server will start on port 8080. You can access it at http://localhost:8080

## Docker Support

To build and run using Docker:

1. Build the image:
```bash
docker build -t go-webserver .
```

2. Run the container:
```bash
docker run -p 8080:8080 go-webserver
```

The server will be accessible at http://localhost:8080 