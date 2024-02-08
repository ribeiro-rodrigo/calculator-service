FROM golang:1.20.7-alpine3.18 as Build

# Set the Current Working Directory inside the container
WORKDIR /usr/local/go/src/server

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download 

# Generate binary
RUN go build -o server calculator_server/main.go 

FROM alpine:3.18

WORKDIR /server

COPY --from=Build /usr/local/go/src/server ./ 

ENTRYPOINT ["./server"]