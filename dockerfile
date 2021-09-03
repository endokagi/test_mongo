FROM golang:1.15.7-alpine3.13 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go build main.go

FROM golang:1.15.7-alpine3.13 AS runtime
WORKDIR /app
COPY --from=builder /app ./
CMD ["/app/main"]