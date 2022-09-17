FROM golang:1.16-alpine as builder
WORKDIR /build/
COPY . .
RUN apk add --no-cache make gcc musl-dev git
RUN go build server.go

FROM alpine:latest
LABEL author=Nathan13888
WORKDIR /app
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/America/Toronto /etc/localtime

ENV USER=rf
ENV UID=12345
ENV GID=23456
RUN addgroup --gid $GID $USER
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN chown -R rf:rf .
RUN chmod -R 755 .
USER rf

COPY --from=builder --chown=rf /build/bin/rf /app/rf

VOLUME /app/.rf
EXPOSE 3000

ENTRYPOINT [ "/app/rf" ]

