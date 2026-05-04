FROM golang:alpine AS builder
WORKDIR /app    
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/app/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /main /main
COPY --from=builder /app/.env /.env
EXPOSE 8080
CMD ["/main"]

# FROM golang:1.24.6-alpine3.22 AS build
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
# go build -o /myapp ./cmd/app

# FROM alpine:3.22 AS run
# COPY --from=build /myapp /myapp
# COPY --from=build /app/.env /.env
# EXPOSE 8080
# CMD ["/myapp"]