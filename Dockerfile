FROM golang:debian as builder
ENV GOPROXY="http://nexus.prod.uci.cu/repository/go-all/"
WORKDIR /app 
COPY . .
RUN tar -xzf vendor.tar
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./cmd

FROM scratch
COPY ./config ./config
COPY --from=builder /app/app /usr/bin/
ENTRYPOINT ["app"]
