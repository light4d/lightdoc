FROM alpine:latest
MAINTAINER light4d
WORKDIR /opt/lightdoc
ADD go/src/lightdoc/lightdoc /opt/lightdoc
ADD vue/dist .
EXPOSE 8000
EXPOSE 8080

ENTRYPOINT ["./opt/lightdoc/lightdoc"]
