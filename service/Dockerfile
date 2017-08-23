FROM alpine

ADD build/ca-certificates.crt /etc/ssl/certs/
ADD dist/epioca-amd64 /

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

# HEALTHCHECK --interval=5s --timeout=5s --retries=3 CMD curl -f http://localhost:8081/status || exit 1

EXPOSE 8080
ENTRYPOINT ["/epioca-amd64"]
