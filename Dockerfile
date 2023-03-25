# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o ./go-avia-crud

EXPOSE 8000

CMD ["./go-avia-crud", "-migratedb", "-seeddb"]
