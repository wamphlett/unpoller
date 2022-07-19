FROM golang:latest as builder

COPY . /build/
WORKDIR /build/
RUN go get && go build -o unpoller

FROM ubuntu

COPY --from=builder /build/unpoller /image
COPY --from=builder /build/examples/up.conf.example /etc/unifi-poller/up.conf

ENV TZ=UTC

ENTRYPOINT [ "/image" ]
