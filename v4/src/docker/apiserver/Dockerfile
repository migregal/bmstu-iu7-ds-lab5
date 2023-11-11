FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

RUN apk add --no-cache make

ADD . .

RUN cd v4/src \
    && make build_apiserver \
    && mv bin/apiserver /bin \
    && mkdir -p /usr/local/etc/apiserver/ \
    && cp apiserver/config-gh.yaml /usr/local/etc/apiserver/config.yaml

FROM alpine

WORKDIR /bin/

COPY --from=builder /bin/apiserver .
COPY --from=builder /usr/local/etc/apiserver/config.yaml /usr/local/etc/apiserver/config.yaml

ENTRYPOINT ["/bin/apiserver"]
