FROM golang:1.20-buster as build

WORKDIR /app

RUN useradd -u 1001 nonroot

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN go build \
    -ldflags "-linkmode external -extldflags -static" \
    -tags netgo \
    -o customer-service ./cmd

###
FROM scratch

ENV APP_ENV=production

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/customer-service customer-service

COPY config.json /config.json

USER nonroot

EXPOSE 3001

CMD ["/customer-service"]