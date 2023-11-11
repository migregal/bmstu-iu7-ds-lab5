FROM golang:alpine as builder
LABEL maintainer="Gregory @migregal Mironov"

RUN mkdir -p /src/backend
WORKDIR /src/backend

RUN apk add --no-cache make

ADD . .

RUN cd v4/src \
    && make build_rating \
    && mv bin/rating /bin \
    && mkdir -p /usr/local/etc/rating/ \
    && cp rating/config-gh.yaml /usr/local/etc/rating/config.yaml


FROM alpine

WORKDIR /bin/

COPY --from=builder /bin/rating .
COPY --from=builder /usr/local/etc/rating/config.yaml /usr/local/etc/rating/config.yaml

ENTRYPOINT ["/bin/rating"]
