FROM golang:1.19-alpine3.16 AS builder

WORKDIR /app

COPY ./merchant .

RUN go build

FROM alpine:3.16

WORKDIR /app

RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache tzdata
ENV TZ="Asia/Jakarta"

COPY --from=builder /app/merchant .
COPY merchant/.env .

EXPOSE 8888

ENTRYPOINT [ "/app/merchant" ]