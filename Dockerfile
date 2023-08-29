# Build stage
FROM golang:1.16 AS build

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o qbook ./cmd/main



# Final stage
FROM scratch
COPY --from=build /app/qbook .
CMD ["./qbook"]
