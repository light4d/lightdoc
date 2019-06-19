FROM light4d/gonode
MAINTAINER light4d
WORKDIR /opt/lightdoc
ADD go /opt/lightdoc/
RUN cd /opt/lightdoc/go/src/lightdoc && go get
ADD vue /opt/lightdoc
RUN cd /opt/lightdoc/vue && npm install && npm run build
EXPOSE 8000 8080
VOLUME ["/var/doc"]
CMD ["/opt/lightdoc/go/bin/lightdoc"]

