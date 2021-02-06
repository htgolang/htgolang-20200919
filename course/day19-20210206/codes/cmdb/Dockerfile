
FROM golang:latest AS builder

WORKDIR /opt/user/
COPY . /opt/user/

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build .

FROM ubuntu:latest

WORKDIR /opt/user/

COPY --from=builder /opt/user/user /opt/user/user
COPY --from=builder /opt/user/conf /opt/user/conf
COPY --from=builder /opt/user/logs /opt/user/logs
COPY --from=builder /opt/user/static /opt/user/static
COPY --from=builder /opt/user/views /opt/user/views
COPY --from=builder /opt/user/tmp /opt/user/tmp


ENTRYPOINT ["/opt/user/user"]