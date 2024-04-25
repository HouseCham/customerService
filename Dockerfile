FROM golang:1.21.6

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go get
RUN go build -o bin

ENTRYPOINT [ "/app/bin" ]