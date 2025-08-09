FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/vicblog
RUN ls -la && echo "Current Directory:" && pwd && echo "Contents of /go/src/github.com/victorzhou123/vicblog" && ls -la /go/src/github.com/victorzhou123/vicblog
RUN cd /go/src/github.com/victorzhou123/vicblog/user-server && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

# setting time zone
RUN apk update && apk add --no-cache tzdata \ 
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \ 
    && echo "Asia/Shanghai" > /etc/timezone \ 
    && rm -rf /var/cache/apk/*

COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/user-server /opt/app
COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/user-server/config/config.yaml /opt/app/config/

ENTRYPOINT ["/opt/app/user-server"]
