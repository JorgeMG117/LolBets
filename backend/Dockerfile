# Use a minimal base image, like Alpine Linux
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

ADD *.mod *.sum ./
RUN go mod download
ADD . .
ADD backend/.env .

# Copy the Go source code into the container
COPY cmd/server/main.go .

# Build the Go program inside the container
RUN go build -o backend main.go

# Set the command to run the executable
CMD ["./backend"]
