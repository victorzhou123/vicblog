FROM golang:1.21 as BUILDER

# build binary
COPY . /go/src/github.com/victorzhou123/vicblog
RUN cd /go/src/github.com/victorzhou123/vicblog/category-server && GO111MODULE=on CGO_ENABLED=0 go build

# copy binary config and utils
FROM alpine:latest
WORKDIR /opt/app/

COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/category-server /opt/app
COPY --from=BUILDER /go/src/github.com/victorzhou123/vicblog/category-server/config/config.yaml /opt/app/config/

ENTRYPOINT ["/opt/app/category-server"]
