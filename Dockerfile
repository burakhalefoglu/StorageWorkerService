FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/storage-worker-service .

EXPOSE 8000

CMD [ "/app/storage-worker-service" ]