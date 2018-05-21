FROM golang:1.9.0 as builder

# RUN go get -u hager-microservices/user-service

WORKDIR /go/src/hager-microservices/user-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/hager-microservices/user-service/user-service .

CMD ["./user-service"]
