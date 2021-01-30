# stage 1
FROM golang:1.15.6 as builder
WORKDIR /app

# fetch dependeicnies first as they're not changing often and will get cached
COPY ./go.mod ./go.sum ./
RUN go mod download

# copy source to working dir of a container
COPY . .

# build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/main.go

# stage 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
ENTRYPOINT ["./server"]
