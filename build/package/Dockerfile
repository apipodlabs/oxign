# Stage 1: Build stage
FROM golang:1.23.4-alpine3.21 AS build

# Set the working directory
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o out cmd/oxign/oxign.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o out cmd/oxignd/oxignd.go

# Stage 2: Final stage
FROM alpine:3.21

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/out/ .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

CMD [ "oxign" ]
