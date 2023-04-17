# !/bin/bash

docker stop api-hosting
docker rm api-hosting

servicio=api-hosting

#PRODUCTION
docker run -d \
  --name $servicio \
  --publish 8081:8081 \
  --network=database_network \
  $servicio
