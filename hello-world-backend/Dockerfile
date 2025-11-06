# ---------- Build stage ----------
FROM golang:1.25.4 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod ./
RUN go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o hello-world-backend .


# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12

# Set working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/hello-world-backend .

# Set environment variables
ENV PORT=8080
ENV GREETING="Hello World!"

# Expose the app port
EXPOSE 8080

# Run the app
ENTRYPOINT ["/app/hello-world-backend"]
