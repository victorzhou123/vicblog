FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/vicblog
RUN cd /go/src/github.com/victorzhou123/vicblog && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

# 安装时区包并设置上海时区
RUN apk update && apk add --no-cache tzdata \ 
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \ 
    && echo "Asia/Shanghai" > /etc/timezone \ 
    && rm -rf /var/cache/apk/*

COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/vicblog /opt/app
COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/config/config.yaml /opt/app/config/

ENTRYPOINT ["/opt/app/vicblog"]
