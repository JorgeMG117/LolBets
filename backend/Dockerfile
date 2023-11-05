# Use a minimal base image, like Alpine Linux
FROM golang:alpine


# Set the working directory inside the container
WORKDIR /app

COPY . .

# Build the Go program inside the container
RUN go build -o backend cmd/server/main.go

EXPOSE 8080

# Set the command to run the executable
CMD ["./backend"]
