# !/bin/bash

docker stop api-marketplace
docker rm api-marketplace

servicio=api-marketplace

#PRODUCTION
docker run -d \
  --name $servicio \
  --publish 8081:8081 \
  --network=database_network \
  $servicio
