# start a golang base image, version 1.8
FROM golang:latest as builder

#switch to our app directory
RUN mkdir -p /go/src/app

WORKDIR /go/src/app

#copy the source files
COPY . /go/src/app

#disable crosscompiling 
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

#build the binary with debug information removed
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# start with a scratch (no layers)
FROM scratch

# copy our static linked library
COPY --from=builder /go/src/app/main /

COPY .env /

ENV MONGODB_CONN=mongodb://mongodb:27017
ENV REDIS_CONN=redis-cache:6379
ENV KAFKA_BROKER=kafka:9092
ENV LOGSTASH_HOST=elk-stack-for-test-logstash:5000

# run it!
CMD ["./main"]