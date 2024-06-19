# GOFILEGO

In this Go-based file sharing application, the structure of the application is defined with the help of several key components and functions. Here's a brief overview:

### Project Setup
First, you initialize the Go module with the command:
```shell
go mod init GOFILEGO
```
This sets up a new Go module named "GOFILEGO." Following this, running `go mod tidy` adds the necessary module requirements and cleans up any unused packages.

### Application Entry Point
The `main.go` file contains the main function, which is the entry point of the application. This function sets up the application routes:

1. `main()` calls the `SetupAppRouter()` function.
2. `SetupAppRouter()` establishes a database connection, configures the Gin router, and sets the application mode.
3. It then initializes the authentication routes by calling `InitAuthRoutes()` with the database connection and the API route group.

### Route Initialization
The `InitAuthRoutes` function is responsible for setting up the authentication routes:

1. It creates a new instance of `RepositoryLogin` using the database connection.
2. Then, it creates a new instance of `ServiceLogin` using the repository instance.
3. Finally, it creates a new instance of `HandlerLogin` using the service instance and sets up the POST `/login` route.

### Call Flow
The application follows a clear call flow:
1. **Handler**: The `loginHandler` is responsible for validating and parsing the request body into a `LoginInput` struct. If any errors occur during this process, it returns an appropriate error response. It then calls the `loginService`.
2. **Service**: The `loginService` takes the `LoginInput`, parses it into a `UserEntity` struct, and calls the `loginRepository`.
3. **Repository**: The `loginRepository`, which has access to the database, performs the required database operations and returns the result.

### Use of Pointers
The application uses pointers for the return types in functions for several reasons:

- **Direct Modification**: By returning a pointer to the `UserEntity` type, it allows direct modification of the object. Any changes made to the returned object will affect the original object. This is useful for maintaining a single shared instance and avoiding unnecessary object copies.
  
- **Avoiding Object Copies**: Returning a value (instead of a pointer) creates a new object each time the function is called. Any modifications made to the returned object will not affect the original object passed as an argument. While this can be useful for creating multiple independent instances, it is generally less efficient due to the additional overhead of object copying.

By using pointers, the application ensures efficient memory usage and maintains consistency in the state of objects across different layers (Handler, Service, Repository).
