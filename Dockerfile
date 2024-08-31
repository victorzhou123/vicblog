FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/vicblog
RUN cd /go/src/github.com/victorzhou123/vicblog && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/vicblog /opt/app
COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/config/config.yaml /opt/app/config

ENTRYPOINT ["/opt/app/vicblog"]
