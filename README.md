# go-api-example

This repository provides a demonstration of a Go-based service with a focus on integrating with Kafka and providing both RESTful and gRPC APIs.

# Prerequisites

Ensure you have Docker and Docker Compose installed as they are essential for setting up and running the app.

# Testing

### Unit Tests

To execute unit tests:

`make test`

### Integration Tests

run integration test (it deals with docker-compose)

`make test-integration`

# Running the Application

## Setting up

First, bring up the necessary services using Docker:

`docker-compose up`

Before running the application, set up Kafka groups:

`docker exec -it kafka \
kafka-consumer-groups \
--bootstrap-server localhost:9094 \
--group go-api-example \
--topic topic-insert-movie \
--reset-offsets \
--to-earliest \
--execute`

Launching the App

`go run cmd/service/app.go`

## Interacting with the Application

### Ingesting a Movie via Kafka

To send a Kafka message that ingests a movie into the database, use the following:

`echo '{"title": "TheMatrix","year": 1999,"duration": 136,"director": "Lana Wachowski, Lilly Wachowski","cast": ["Keanu Reeves", "Laurence Fishburne"],"genre": ["Science Fiction", "Action"],"synopsis": "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.","box_office_revenue": 463517383}' | docker exec -i kafka kafka-console-producer --broker-list localhost:9094 --topic topic-insert-movie`

#Fetching a Movie

## getting movie using /movies/{name} endpoint in RESTful API

`curl --location 'localhost:8001/movies/TheMatrix'`

## getting movie using GetMovie endpoint in gRPC API

`grpcurl -plaintext -d '{"title": "TheMatrix"}' localhost:50051 movieapi.MovieService/GetMovie`