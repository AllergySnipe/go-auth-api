# Go Auth API

This is a Go-based REST API for user authentication, which includes JWT-based authentication and user sign-up/login functionality. The application connects to a PostgreSQL database (Tembo), and allows you to generate, refresh, and revoke JWT tokens for authorized access.

## Features

- User registration (sign-up)
- User login with password hashing (bcrypt)
- JWT token generation for secure authentication
- Token refresh and revocation
- Protected API routes using JWT middleware

## Requirements

- Go 1.18 or later
- Postman for testing

## Setup

### 1. Clone the repository

Clone the repository to your local machine:

```bash
git clone https://github.com/AllergySnipe/go-auth-api.git
cd go-auth-api
```


### 2. Build the Project

To build the Go application, run:

```bash
go build -o go-auth-api main.go
```

### 3. Running the Server

To run the server, execute the following command:

```bash
./start.sh
```

This script will:

1. Set the necessary environment variables for connecting to the Tembo database and JWT secret.
2. Build the Go application if it's not already built.
3. Start the server.

### 4. Testing the API

You can use **Postman** or **curl** to test the API:

#### Sign Up:
- **POST** `/signup`
- Request Body: 
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```
  curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "yourpassword"
  }'

#### Sign In (Login):
- **POST** `/signin`
- Request Body: 
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```
  curl -X POST http://localhost:8080/signin \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "yourpassword"
  }'

If successful, the response will include the JWT token in the `Authorization` header. You can then use this token to make requests to protected endpoints.

#### Protected API (Auth Middleware):
- **GET** `/auth/protectedroute`
- Requires JWT token in the `Authorization` header (by default passed as the time of signing in if using Postman).
- If not signed in, does not grant access

Example using `curl` to access a protected route:

```bash
curl -H "Authorization: Bearer your-jwt-token" http://localhost:8080/auth/protectedroute
```

### 5. Refreshing the Token

If your token has expired, you can refresh it by calling the **POST** `/auth/refresh` endpoint, which will return a new token.
curl -X POST http://localhost:8080/auth/refresh \
  -H "Authorization: Bearer your-jwt-token"

### 6. Revoke Token

To revoke a token (invalidate the JWT), call the **POST** `/auth/revoke` endpoint.
curl -X POST http://localhost:8080/auth/revoke \
  -H "Authorization: Bearer your-jwt-token"

---
