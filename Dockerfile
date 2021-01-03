FROM golang:1.15.6-alpine as builder

RUN apk update && apk add upx

RUN mkdir /app
WORKDIR /app

ADD go.mod go.sum /app/
RUN go mod download

# Copy sources
ADD . /app/

# Compile
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Compress
RUN upx wiki

################################## Runner
FROM scratch as runner

WORKDIR /app

COPY --from=builder /app/wiki /app/
EXPOSE 3000
CMD ["/app/wiki","-port=3000"]
