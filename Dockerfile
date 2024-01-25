# Stage 1: Build Stage
FROM golang:1.19 as build

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/server.go

FROM scratch

WORKDIR /app
COPY .env .
COPY --from=build /app/server .
EXPOSE 8000
CMD ["./server"]