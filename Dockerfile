FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/
ENV APP_ENV=prod

CMD ["./main"]