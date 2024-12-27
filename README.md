![TASK](https://github.com/user-attachments/assets/c9dc71ab-55fd-4e0c-b5c3-9b20e04ee651)

# Project task-tracker
The issue tracking system is designed to streamline the management of project-related issues by providing a centralized platform for logging, monitoring, and resolving tasks.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

To Run the services, Make sure you installed Make, Docker. Just go inside the service directory and give the following make commands.

Eg; cd task-tracker-backend; make run 

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```
Create DB container
```bash
make docker-run
```

Shutdown DB Container
```bash
make docker-down
```

DB Integrations Test:
```bash
make itest
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
