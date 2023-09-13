# go-api-example

 echo "Your message here" | docker exec -i kafka kafka-console-producer \
  --broker-list localhost:9094 \
  --topic topic-insert-movie

docker exec -it kafka \
kafka-consumer-groups \
--bootstrap-server localhost:9094 \
--group go-api-example \
--topic topic-insert-movie \
--reset-offsets \
--to-earliest \
--execute