FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/vicblog
RUN cd /go/src/github.com/victorzhou123/vicblog/category-server && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

# setting time zone
RUN apk update && apk add --no-cache tzdata \ 
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \ 
    && echo "Asia/Shanghai" > /etc/timezone \ 
    && rm -rf /var/cache/apk/*

COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/category-server /opt/app
COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/category-server/config/config.yaml /opt/app/config/

ENTRYPOINT ["/opt/app/category-server"]
