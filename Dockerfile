FROM golang
MAINTAINER Ragnar B. Johannsson <ragnar@igo.is>

ADD . /go/src/github.com/ragnar-johannsson/lto
RUN go install github.com/ragnar-johannsson/lto

ENTRYPOINT /go/bin/lto

EXPOSE 3000
