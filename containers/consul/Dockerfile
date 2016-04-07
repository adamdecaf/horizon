FROM alpine:3.3

RUN apk --update upgrade && \
    apk add curl ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

ADD https://releases.hashicorp.com/consul/0.6.4/consul_0.6.4_linux_amd64.zip /opt/consul.zip
RUN cd /bin && unzip /opt/consul.zip && chmod +x /bin/consul && rm /opt/consul.zip

# ADD https://releases.hashicorp.com/consul/0.6.4/consul_0.6.4_web_ui.zip /opt/webui.zip
# RUN cd /opt && unzip /opt/webui.zip && mv dist /ui && rm /opt/webui.zip

VOLUME /opt/consul/

ENV CONSUL_DATA_DIR /opt/consul/

ADD start-consul.sh /bin/start-consul
RUN chmod u+x /bin/start-consul
CMD ["/bin/start-consul"]
