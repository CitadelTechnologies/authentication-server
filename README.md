# Citadel Authentication Server

This repository contains an authentication server written in Go for micro-services architectures.

The server shares session data through Redis and exposes a REST API for registration, login and services.

## Setup

To use the image, pull it from Docker Hub:

```
docker pull citadeltechnologies/authentication-server
```

Then you can run the container with Docker or Docker Compose.

The server listens on port 80, you can expose it externally with Docker port mapping.

The image needs several environment variables, listed below with default values.

```
MYSQL_HOST=127.0.0.1
MYSQL_USER=root
MYSQL_PASSWORD=my_incredible_password
MYSQL_DBNAME=my_wonderful_database
MYSQL_TEST_DBNAME=my_wonderful_database # useful only if you want to run the tests
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASSWORD=redis_password
```

## API documentation

You can read the API documentation here:

* [Registration](./doc/registration.md)
