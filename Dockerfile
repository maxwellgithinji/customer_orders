# Start from golang base image
FROM golang:1.16.2

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app

# We copy everything in the root directory
# into our /app directory
ADD . /app
RUN touch /app/.env

# Set the current working directory inside the container 
WORKDIR /app

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

RUN go build -o main

CMD ["/app/main"]
