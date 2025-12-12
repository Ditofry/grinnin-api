# Grinnin Golang API

A basic HTTP REST API built with Go using only standard library packages.

## Features

- Single `/hello` endpoint supporting GET and POST methods
- GET: Returns `{"message": "Hello World"}`
- POST: Accepts JSON body with `name` field and returns `{"message": "Hello {name}"}`

## Running the Server

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

### GET /hello

Returns a simple "Hello World" message.

**Response:**
```json
{
  "message": "Hello World"
}
```

### POST /hello

Accepts a name in the request body and returns a personalized greeting.

**Request Body:**
```json
{
  "name": "Brandon"
}
```

**Response:**
```json
{
  "message": "Hello Brandon"
}
```

## Example Usage

### GET Request
```bash
curl http://localhost:8080/hello
```

### POST Request
```bash
curl -X POST http://localhost:8080/hello \
  -H "Content-Type: application/json" \
  -d '{"name":"Brandon"}'
```

