FROM golang:alpine AS builder

WORKDIR /builder

RUN apk add --no-cache gcc musl-dev

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=1
RUN go build -o ./cmd/app/main ./cmd/app/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /builder/cmd/app/main ./cmd/app/main
COPY --from=builder /builder/migrations/ ./migrations

EXPOSE 8080

CMD ["./cmd/app/main"]
