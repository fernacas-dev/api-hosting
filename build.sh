# !/bin/bash

servicio=api-hosting

echo "Empaquetando $servicio"

#ruta=$PWD
ruta=/home/cloud/golang_projects/api-hosting

#docker stop $servicio
#docker rm $servicio

#docker rmi $servicio

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

#echo "BUILD APP"
#go build -a -installsuffix cgo -o app ./cmd

cd $ruta

docker rmi $servicio

echo "BUILD IMAGE"
docker build -t $servicio .

#rm app

#docker build -t registry.gitlab-vcl.xutil.net:5000/devops/$servicio .
#docker push registry.gitlab-vcl.xutil.net:5000/devops/$servicio

#docker run -p 3100:3000 --name=$servicio --restart=always -d $servicio
