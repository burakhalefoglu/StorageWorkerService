FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /storage-worker-service .

RUN cd ..
RUN rm -rf app
RUN mkdir app

EXPOSE 8000

CMD [ "/storage-worker-service" ]