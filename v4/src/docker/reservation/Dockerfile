FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

RUN apk add --no-cache make

ADD . .

RUN cd v4/src \
    && make build_reservation \
    && mv bin/reservation /bin \
    && mkdir -p /usr/local/etc/reservation/ \
    && cp reservation/config-gh.yaml /usr/local/etc/reservation/config.yaml


FROM alpine

WORKDIR /bin/

COPY --from=builder /bin/reservation .
COPY --from=builder /usr/local/etc/reservation/config.yaml /usr/local/etc/reservation/config.yaml

ENTRYPOINT ["/bin/reservation"]
