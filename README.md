# goStorage

`goStorage` is a simple web service built with Go that implements a REST API for interacting with a key-value store. It supports operations for adding, retrieving, and deleting values by key. Additionally, the service logs all transactions and can restore its state upon the next startup.

## Features

- **PUT** `/v1/{key}`: Stores a value associated with the specified key. The value is provided in the request body.
- **GET** `/v1/{key}`: Retrieves the value associated with the specified key.
- **DELETE** `/v1/{key}`: Deletes the value associated with the specified key.
- **Transaction Log**: All operations are logged, allowing the service to restore its state from the log file when restarted.

## Installation

### Prerequisites

Make sure you have Go installed (version 1.16 or higher). You can download it from the [official site](https://golang.org/dl/).

### Clone the Repository

git clone https://github.com/your_username/goStorage.git
cd goStorage


### Install Dependencies

go mod tidy

## Running the Application

To start the service, run the following command:

go run cmd/server/main.go


The server will listen on port `8080`.

## Usage Examples

### Adding a Value (PUT)

curl -X PUT -d "value" http://localhost:8080/v1/key


### Retrieving a Value (GET)

curl http://localhost:8080/v1/key


### Deleting a Value (DELETE)

curl -X DELETE http://localhost:8080/v1/key


## Transaction Log

All operations are logged to a file. Upon the next startup, the service restores its state from this log file.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more details.

---

Thank you for using `goStorage`! If you have any questions or suggestions, feel free to reach out.