# Go Redis Example

This is a simple example project that demonstrates how to connect to a Redis database using Go and the `go-redis` library.

## Prerequisites

Before running this project, make sure you have the following prerequisites:

- Go installed on your machine. You can download and install Go from the official website: https://go.dev/dl/
- Docker installed on your machine. You can download and install Docker from the official website: https://www.docker.com/get-started

## Getting Started

Follow the steps below to set up and run the project:

1. Clone the repository:
```bash
$ git clone https://github.com/lucasgrvarela/go-redis-example.git
```
2. Change into the project directory:

```bash
$ cd go-redis-example
```

3. Start a Redis server using Docker:
```bash
$ docker run -d -p 6379:6379 --name redis-server redis
```

4. Run the Go application:
```bash
$ go run *.go
Key set successfully!
Key: 'mykey' retrieved, value: 'myvalue'
```

5. Test manually everything worked as expected
```bash
$ docker exec -it redis-server bash
$ redis-cli # run inside the container
127.0.0.1:6379> GET mykey
"myvalue"
```

# Configuration
By default, the Go application connects to Redis running on localhost:6379. If your Redis server is running on a different address or port, you can modify the main.go file and update the Redis server address in the NewRedisClient function.

# Contributing
Contributions are welcome! If you find any issues with the project or want to extend it further, feel free to open an issue or submit a pull request.
