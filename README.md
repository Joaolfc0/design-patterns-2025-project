# Secret Santa Backend for Design Patterns UBB

## Author Name

- *Joao Luiz Figueiredo Cerqueira*

## System Explanation

This system consists of an API for a Secret Santa application called *Secret Santa*. The API allows managing groups and participants, performing automatic draws to define the Secret Santa pairs, and consulting information related to groups and participants. The main features include creating, updating, deleting groups, adding participants, generating matches automatically, and securely and efficiently retrieving information. This API uses 5 design pattern implementations, as shown in the documentation PDF.

### Main Implemented Routes

- *POST /group* - Creates a new group.
- *GET /group/:id* - Retrieves the details of a specific group.
- *PUT /group/:id* - Updates an existing group.
- *DELETE /group/:id* - Deletes an existing group.
- *POST /group/:id/add-participant* - Adds a participant to a group.
- *POST /group/:id/match-participants* - Performs the group's participants draw.
- *GET /group/:id/my-match* - Retrieves the assigned Secret Santa pair for a participant.
- *GET /group* - Retrieves all registered groups.

The route structure was configured using the *Gin* framework, allowing clear and efficient organization of HTTP requests.

## Explanation of the Technologies Used

- *Go Lang:* The main language used to develop the API due to its efficiency and robustness.

- *Gin Framework:* Framework used to build fast and minimalist RESTful APIs.

- *Uber Dependency Injection:* Library used for dependency injection, facilitating system maintenance and scalability.

- *MongoDB:* NoSQL database used for flexible and scalable data storage.

- *Dockerfile and Docker Compose:* Used to create Docker images and configure the development and production environments.

- *Go Mock:* Used for creating mocks and simulating dependencies in unit tests.

- *Swagger with Swaggo:* Used to automatically generate API documentation.

- *"go.mongodb.org/mongo-driver/mongo/integration/mtest":* Library used to mock MongoDB operations during unit tests.

- *Validation with "github.com/invopop/validation":* Library for validating DTOs and data structures received by the API.

- *Custom Errors:* Custom implementation for managing application-specific errors, facilitating error handling and debugging.

- *CI for Unit Tests:* Continuous integration configured to automatically run unit tests upon each update to the repository, ensuring the system’s quality and stability.

This set of technologies was chosen to ensure a robust, scalable, and secure API, following the best practices of modern software development.

---

## How to Run

You can run the project using Docker Compose or run the Go API locally while using Docker for MongoDB.

### Option 1: Run Entirely with Docker Compose

```bash
go mod tidy
docker compose up
```

After the application is running, access the Swagger documentation at:

```
http://localhost:8080/secret-santa/swagger/index.html#/
```

This will bring up the full API interface.

---

### Option 2: Run API Locally and MongoDB in Docker

If you want to run only the database in Docker and run the Go API on your machine:

1. Start only MongoDB:

```bash
docker compose up mongo
```

2. Set your local `.env` file or export this environment variable:

```bash
export MONGO_URI=mongodb://localhost:27017/secret-santa
```

3. Then run the Go API locally:

```bash
go mod tidy
go run main.go
```

4. Open Swagger:

```
http://localhost:8080/secret-santa/swagger/index.html#/
```
