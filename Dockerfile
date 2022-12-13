FROM golang:1.19-alpine

WORKDIR /app
COPY . .

RUN go mod download

CMD go run file-server/cmd/server
