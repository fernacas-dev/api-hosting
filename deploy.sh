# !/bin/bash

docker stop api-hosting
docker rm api-hosting

servicio=api-hosting

#PRODUCTION
docker run -d \
  --name $servicio \
  --publish 8081:8080 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  $servicio
