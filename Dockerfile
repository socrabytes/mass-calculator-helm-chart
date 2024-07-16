# Use a multi-stage build to keep the final image size small
FROM golang:1.19-alpine AS builder
# Set the working directory inside the container
WORKDIR /app
# Copy the Go module files and download dependencies
COPY src/go.mod ./
RUN go mod download
# Copy the source code and build the application
COPY src/ ./
RUN go build -o main .
# Use a minimal image for the final stage
FROM alpine:latest
# Create a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# Set the working directory inside the container
WORKDIR /app
# Copy the built binary from the builder stage
COPY --from=builder /app/main .
# Set the user to the non-root user
USER appuser
# Expose the port (optional, for documentation purposes)
EXPOSE 8080
# Set the environment variable for the port
ENV PORT=8080
# Run the application
ENTRYPOINT ["sh", "-c", "./main $PORT"]
