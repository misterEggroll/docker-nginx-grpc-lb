FROM golang:1.13-alpine as builder

RUN mkdir -p $GOPATH/src/demo_grpc/ && \
    mkdir -p $GOPATH/src/demo_grpc/build/

WORKDIR $GOPATH/src/demo_grpc/

COPY . .

ENV GO111MODULE on

RUN go build ./...

# Create executable file for http server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o register registrar/main.go

FROM alpine

# Copy the register executable to the bin from where it can be used
COPY --from=builder /go/src/demo_grpc/register /bin/

ADD ../config/supervisord /etc/

ENV CONSUL_VERSION=1.9.1 \
    CONSUL_DOMAIN=consul \
    CONSUL_DATA_DIR=/data/consul \
    CONSUL_CONFIG_DIR=/etc/consul/conf.d/bootstrap \
    CONSUL_SERVER_NAME=consul \
    CONSUL_DC=dc1 \
    CONSUL_CLIENT=0.0.0.0 \
    CONSUL_RETRY_INTERVAL=5s

# Download and install Consul
RUN apk upgrade --update --no-cache && \
    apk add --update --no-cache curl util-linux supervisor && \
    mkdir -p /var/log/supervisor/ && \
    mkdir -p /etc/supervisor/ && \
    curl -sSLo /tmp/consul.zip https://releases.hashicorp.com/consul/{$CONSUL_VERSION}/consul_${CONSUL_VERSION}_linux_amd64.zip && \
    unzip -d /bin /tmp/consul.zip && \
    rm /tmp/consul.zip && \
    apk del curl && \
    addgroup consul && \
    adduser -D -g "" -s /bin/sh -G consul consul && \
    mkdir -p /data/consul && \
    chown -R consul:consul /data/consul && \
    rm -rf /tmp/* /var/cache/apk/* && \
    mkdir /etc/consul.d/

# Add the files
COPY ../config/consul /
COPY ../config/consul /etc/supervisor/conf.d/consul.conf
COPY ../config/supervisord /etc/supervisor/conf.d/register.conf
COPY ../config/supervisord /etc/supervisor/supervisord.conf


VOLUME ["/data/consul"]

# Same exposed ports than consul
EXPOSE 8300 8301 8301/udp 8302 8302/udp 8400 8500 8600 8600/udp 53 53/udp

ENTRYPOINT ["supervisord", "-c", "/etc/supervisor/supervisord.conf"]
# Command to run: docker run -p 8500:8500 -e CONSUL_SERVER=127.0.0.1 -d -t base_img_test
