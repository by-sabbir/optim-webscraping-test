FROM golang:1.21 as build

ENV CGO_ENABLED 0
ARG BUILD_REF

COPY . /service/app

WORKDIR /service/app
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM alpine:3

ARG BUILD_DATE
ARG BUILD_REF

RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -h / -G app -S app


COPY --from=build --chown=app:app /service/app/optim-webscraping-test /usr/local/bin/scraper

WORKDIR /
USER app

LABEL org.opencontainers.image.created="${BUILD_DATE}}"\
      org.opencontainers.image.version="${BUILD_REF}"\
      org.opencontainers.image.title="app"