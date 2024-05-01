# CustomerService: Microservice Connected to Cerebro API Gateway

CustomerService is a microservice written in Golang, designed to be connected to the Cerebro API Gateway. It provides specific functionalities and services to the overall system.

## Features

- Specific functionalities for customer-related operations
- Integration with Cerebro API Gateway via gRPC connections
- Scalable architecture for handling customer-related requests

## Structure
```
customerService/
    ├── api/
    │    └── core/
    │         └── grpc/
    ├── cmd/
    │    └── main.go
    ├── internal/
    │    ├── config/
    │    ├── handler/
    │    ├── model/
    │    ├── repository/
    │    ├── validator/
    │    └── log/
    ├── pkg/
    │    └── utils/
    ├── scripts/
    ├── .dockerignore
    ├── config.json
    ├── go.mod
    ├── go.sum
    ├── README.md
    └── Dockerfile
```


- `api/`: Contains API-related code.
    - `core/`: Core functionality of the API.
        - `grpc/`: gRPC related code.

- `cmd/`: Main application entry point.
    - `main.go`: Main entry point for the application.

- `internal/`: Internal package containing application-specific code.
    - `config/`: Configuration related code.
    - `handler/`: Request handler functions.
    - `model/`: Data models used by the application.
    - `repository/`: Data access layer for interacting with databases.
    - `validator/`: Validation logic for incoming requests.
    - `log/`: Logging functionality.

- `pkg/`: Package directory containing reusable code.
    - `utils/`: Utility functions and libraries.

- `scripts/`: Directory containing scripts for building, testing, and running the application.

- `.dockerignore`: File specifying paths to exclude when building Docker images.

- `config.json`: Configuration file for the application.

- `go.mod`, `go.sum`: Go module files specifying dependencies for the project.

- `README.md`: Markdown file containing project information and instructions.

- `Dockerfile`: File used to build a Docker image for the project.

## Usage

The project can be built and run using Docker containers. Follow the instructions provided in the README.md file to build and run the application.

Feel free to reach out if you have any questions or need further assistance!


## Frameworks & Libraries

The project relies on the following frameworks and libraries:

- [Validator v10](https://github.com/go-playground/validator/v10) v10.19.0
- [Logrus](https://github.com/sirupsen/logrus) v1.9.3
- [Viper](https://github.com/spf13/viper) v1.18.2
- [Crypto](https://golang.org/x/crypto) v0.22.0
- [gRPC](https://google.golang.org/grpc) v1.59.0
- [Protocol Buffers](https://github.com/protocolbuffers/protobuf) v1.33.0
- [GORM Postgres Driver](https://gorm.io/docs/postgres.html) v1.5.7


## Usage

### Building Docker Images

To build the Docker images for CustomerService, follow these steps:

1. Clone the CustomerService repository:

```bash
  git clone https://github.com/HouseCham/customerService.git
```

2. Navigate to the project directory:
```bash
cd customerService
```

3. Build the docker image using the provided Dockerfile:
```bash
docker build -t customer-service:[version] .
```

### Running Docker Containers

Once you have built the Docker image, you can run containers for CustomerService using the following commands:

```bash
sudo docker run -d \
--name customer-service \
-p [host_port]:[container_port] \
-it \
--restart unless-stopped \
customer-service:[version]
```

Replace [host_port] and [container_port] with the desired port mapping. Remember to adjust the configuration within the config.json file accordingly.

### Interacting with the Running Container

If you choose to run the container interactively (-it tag), you can interact with the container's CLI. Upon successful startup, you'll see log messages indicating the server is running.

