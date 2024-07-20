FROM scratch
MAINTAINER Fatih DEMIR <fath.demmr@gmail.com>
ENV PORT=8080
COPY api api
ADD ../../ca-certificates.crt /etc/ssl/certs/
EXPOSE $PORT
COPY ../../zoneinfo.zip /
COPY ../../.env .env
ENV ZONEINFO=/zoneinfo.zip

ENTRYPOINT ["./api"]