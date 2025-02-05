FROM golang:1.23-alpine AS builder

ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}

WORKDIR /usr/src/app

# setup dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# build application
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -o ${SERVICE_NAME} ./cmd/main.go

FROM alpine:latest  

ARG SERVICE_NAME
ENV SERVICE_NAME=${SERVICE_NAME}

WORKDIR /root/
EXPOSE 23456

# export binary from build into final image
COPY --from=builder /usr/src/app/$SERVICE_NAME ./app

CMD ["./app"]
