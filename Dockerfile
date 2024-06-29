# Stage 1: Build the executable library
FROM golang:1.22.3 AS builder

# Set the working directory inside the container
WORKDIR /go/src/github.com/space-code/go-auth

# Copy the entire current direcroty into the container's working directory
COPY . .

# Make dependencies up to date
RUN go mod tidy

# Build the executable binary
RUN GOOS=linux go build -o main cmd/go-auth/main.go

# Stage 2: Create a mininal production image
FROM scratch AS production

# Copy the executable binary from the builder stage
COPY --from=builder /go/src/github.com/space-code/go-auth .

# Set the command to run the container starts
CMD ["./main"]
