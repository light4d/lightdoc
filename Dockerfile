FROM light4d/gonode
MAINTAINER light4d

ADD go /opt/lightdoc/go
RUN cd /opt/lightdoc/go/src/lightdoc &&export GOPATH=/opt/lightdoc/go && go get &&mv /opt/lightdoc/go/bin/lightdoc /opt/lightdoc/lightdoc &&rm -rf  /opt/lightdoc/go
ADD vue /opt/lightdoc/vue
RUN cd /opt/lightdoc/vue && npm install && npm run build &&mv /opt/lightdoc/vue/dist /opt/lightdoc/dist &&rm -rf /opt/lightdoc/vue &&tree
EXPOSE 8050 8051
WORKDIR /opt/lightdoc
VOLUME ["/var/doc"]
CMD [".lightdoc"]

