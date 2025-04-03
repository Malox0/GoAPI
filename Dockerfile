FROM golang:1.24

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./
COPY src/GoAPI/main.go ./
RUN go build -o bin/docker-demo

EXPOSE 8080

CMD ["bin/docker-demo"]