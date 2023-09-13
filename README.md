# go-api-example



# running app

## set up

`docker-compose up`

Set up kafka groups before running the app

`docker exec -it kafka \
kafka-consumer-groups \
--bootstrap-server localhost:9094 \
--group go-api-example \
--topic topic-insert-movie \
--reset-offsets \
--to-earliest \
--execute`


`go run cmd/service/app.go`

## testing the app

Sending kafka message to ingest movie to the database

`echo '{"title": "TheMatrix","year": 1999,"duration": 136,"director": "Lana Wachowski, Lilly Wachowski","cast": ["Keanu Reeves", "Laurence Fishburne"],"genre": ["Science Fiction", "Action"],"synopsis": "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.","box_office_revenue": 463517383}' | docker exec -i kafka kafka-console-producer --broker-list localhost:9094 --topic topic-insert-movie`

## getting movie using /movies/{name} endpoint in RESTful API

`curl --location 'localhost:8001/movies/TheMatrix'`

## getting movie using GetMovie endpoint in gRPC API

`grpcurl -plaintext -d '{"title": "TheMatrix"}' localhost:50051 movieapi.MovieService/GetMovie`