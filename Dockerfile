FROM golang:1.24.10-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o avito_test ./cmd/app

FROM alpine:latest
COPY --from=builder /app/avito_test /avito_test
CMD [ "/avito_test" ]