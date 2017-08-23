FROM golang:alpine

ADD . /go/src/github.com/italolelis/epioca/service

WORKDIR /go/src/github.com/italolelis/epioca/service

ADD build/ca-certificates.crt /etc/ssl/certs/

RUN apk add --update git curl make && \
    rm -rf /var/cache/apk/*

# HEALTHCHECK --interval=5s --timeout=5s --retries=3 CMD curl -f http://localhost:8081/status || exit 1

RUN make deps

EXPOSE 8080
