# !/bin/bash

docker stop api-hosting
docker rm api-hosting

servicio=api-hosting

#PRODUCTION
docker run -d \
  --name $servicio \
  --publish 8081:8081 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --network=appwrite_runtimes \
  $servicio
