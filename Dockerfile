FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o ecommerce ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/ecommerce /

EXPOSE 8000

ENTRYPOINT [ "/ecommerce", "config/local.yaml" ]
