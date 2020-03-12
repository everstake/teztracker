FROM golang:latest AS build-env

RUN mkdir /app
WORKDIR /app
COPY . .
ENV CGO_ENABLED 0
RUN go build

FROM alpine:latest
RUN apk add ca-certificates

RUN mkdir -p repos/migrations

COPY --from=build-env /app/repos/migrations repos/migrations
COPY --from=build-env /app/teztracker /

CMD ["/teztracker"]
