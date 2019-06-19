FROM light4d/gonode
MAINTAINER light4d
WORKDIR /opt/lightdoc
ADD go /opt/lightdoc/go
RUN tree && cd /opt/lightdoc/go/src/lightdoc &&export GOPATH=/opt/lightdoc/go && go get
ADD vue /opt/lightdoc/vue
RUN cd /opt/lightdoc/vue && npm install && npm run build
EXPOSE 8000 8080
VOLUME ["/var/doc"]
CMD ["/opt/lightdoc/go/bin/lightdoc"]

