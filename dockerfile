FROM ubuntu:latest
MAINTAINER light4d
WORKDIR /opt/lightdoc
ADD go/src/lightdoc/lightdoc /opt/lightdoc/
ADD vue/dist /opt/lightdoc/dist
EXPOSE 8000 8080
VOLUME ["/var/doc"]
CMD ["/opt/lightdoc/lightdoc"]


