# stage build
FROM golang:alpine3.16 as build
WORKDIR /
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go .
RUN go build -o /app

# stage release
FROM alpine:latest
WORKDIR /
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT [ "/app" ]