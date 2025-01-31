FROM golang:1.19-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify && \
    go mod tidy
COPY . .
RUN go build -o main ./cmd/main.go

# Final stage: create a smaller image with only the built binary
FROM --platform=$BUILDPLATFORM alpine:3.20 AS final

# Set working directory
WORKDIR /root/

# Copy the configuration file
COPY ./config.json .

# Copy the compiled binary from the build stage
COPY --from=build /app/main .

# Set the default command to run the application
CMD ["./main"]