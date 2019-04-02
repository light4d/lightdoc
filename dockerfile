FROM alpine:latest
MAINTAINER light4d
WORKDIR /opt/lightdoc
ADD src/lightdoc/lightdoc /opt/lightdoc
EXPOSE 8000
EXPOSE 80

ENTRYPOINT ["./opt/lightdoc/lightdoc"]
