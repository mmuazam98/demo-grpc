# Build a simple gRPC service with Go

This repository demonstrates a simple implementation of a gRPC-based microservice for User CRUD (Create, Read, Update, Delete) operations. We have used Go and Cobra to manage both server and client commands in the same binary.

## Key Components

1. Cobra for Command Management:

   - We use Cobra to organize the project into commands.
   - Two main commands:
     - `server`: Starts the gRPC server.
     - `client`: Sends requests to the gRPC server (like creating or retrieving users).

2. gRPC User Service:

   - The gRPC service provides basic CRUD functionality for user entities.
   - `service/user_service.go` contains the logic for the following operations:
     - `CreateUser`: Adds a new user.
     - `DeleteUser`: Deletes a user by ID.
     - `GetUser`: Retrieves user information by ID.
     - `UpdateUser`: Updates user details.

3. Proto Definitions:
   - The gRPC service is based on a `user.proto` file (located in the `user/` directory).
   - This file defines the `UserService` and the request/response types for the operations.

## How to run

1. Install Dependencies: Make sure you have Go installed and then run:

```bash
go mod tidy
```

2. Run the gRPC Server: Use the server command to start the gRPC server.

```bash
go run main.go server
```

3. Run the gRPC Client: Use the client command to send requests to the server. For example, to create a user:

```bash
go run main.go client
```

Note: To generate gRPC code, ensure you generate the necessary gRPC files from the proto:

```bash
make generate_grpc_code
```
