# Device Manager API

**Device Manager** is a RESTful API developed in **Go** for managing devices. The project uses **Clean Architecture**, [Huma](https://github.com/danielgtaylor/huma) for building modern APIs and [GORM](https://gorm.io/) as an ORM for database interactions.

## Features

- [OpenAPI](https://www.openapis.org/what-is-openapi)-compatible API, with documentation
- Add, list, update, and delete devices
- Query devices by unique identifier (ID)
- Input data validation with Huma

---

## Requirements

- **Go 1.23+**
- **Docker** and **Docker Compose** (for containerized execution)

---

## Installation and Local Execution

1. Clone the repository:
   ```bash
   git clone https://github.com/mendelgusmao/device-manager
   cd device-manager
   ```

2. Install dependencies:
   ```bash
   make setup
   ```

3. Run the application: 
    ```bash
    make run
    ```

4. Access the application at [http://localhost:8080](http://localhost:8080).

## Running unit and integration tests: 

* Both unit and integration tests can be run with this command:
    ```bash
    make test
    ```

---

## Running with Docker

1. Build the Docker image:
   ```bash
   make docker/build-image
   ```

2. Run the container:
   ```bash
   make docker/run
   ```

3. Access the application at [http://localhost:8080](http://localhost:8080).

---

## Live Demo

The API is available for testing at:

> [https://device-manager.fawn-beaver.ts.net](https://device-manager.fawn-beaver.ts.net)

---

## API Documentation - OpenAPI Compatible

### Local
> [http://localhost:8080/docs](http://localhost:8080/docs)

### Live
> [https://device-manager.fawn-beaver.ts.net/docs](https://device-manager.fawn-beaver.ts.net/docs)

---

## Technical debts

- I found too difficult to write repository tests using gorm mocks
- Container is running as root
- It lacks proper grouping for API versioning
- All unit tests cover only happy paths
- Integration test could be more DRY
- Not linting
- CORS on the REST layer
- Proper logging and observability
- Is kinda complex for a CRUD

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

