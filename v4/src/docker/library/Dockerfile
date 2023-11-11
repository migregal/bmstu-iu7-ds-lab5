FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

RUN apk add --no-cache make

ADD . .

RUN cd v4/src \
    && make build_library \
    && mv bin/library /bin \
    && mkdir -p /usr/local/etc/library/ \
    && cp library/config-gh.yaml /usr/local/etc/library/config.yaml


FROM alpine

WORKDIR /bin/

COPY --from=builder /bin/library .
COPY --from=builder /usr/local/etc/library/config.yaml /usr/local/etc/library/config.yaml

ENTRYPOINT ["/bin/library"]
