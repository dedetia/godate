FROM golang:1.22-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add make

WORKDIR /app

COPY . .

RUN make build

FROM alpine

RUN apk --update --no-cache add ca-certificates && \
	addgroup -S godate && adduser -S -g godate godate

RUN mkdir -p /config
RUN chown -R godate:godate /config

USER godate
COPY --from=builder /app/godate /
COPY --from=builder /app/config/. /config/.

CMD ["./godate"]

EXPOSE 8000