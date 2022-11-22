FROM golang:latest

# Install git

# RUN apk update && apk add --no-cache git

RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Set the Current Working Directory inside the container

WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container

COPY . .

# Download all the dependencies

RUN go mod download

# Build the Go app

RUN go build -o main cmd/main.go

# Expose port 8080 to the outside world

EXPOSE 8080

# Command to run the executable

# CMD ["./main"]