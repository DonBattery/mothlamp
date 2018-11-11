FROM golang AS builder
WORKDIR /go/src/github.com/donbattery/mothlamp/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mothlamp . 

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/github.com/donbattery/mothlamp/mothlamp .
COPY ./mothlamp.yaml .

CMD ["./mothlamp"]