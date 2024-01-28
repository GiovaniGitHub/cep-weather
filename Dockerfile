# Stage 1: Build Stage
FROM golang:1.19-alpine as build

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/server.go

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /app
COPY .env .
COPY --from=build /app/server .
EXPOSE 8000
CMD ["./server"]