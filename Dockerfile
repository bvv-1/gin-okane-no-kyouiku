FROM golang:1.18-alpine3.15 AS go

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /app/main

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
