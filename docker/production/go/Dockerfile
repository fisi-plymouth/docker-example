FROM golang:1.21-alpine as BUILD_STEP

COPY src/go /go-app

WORKDIR /go-app

RUN go build -o compiled-go-app

RUN chmod +x compiled-go-app


FROM alpine:latest

WORKDIR /go-app

COPY --from=BUILD_STEP /go-app/compiled-go-app .

CMD ["./compiled-go-app"]
