FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o program main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/ .

CMD ["/app/program"]