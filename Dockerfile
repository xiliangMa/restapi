FROM golang:1.11
MAINTAINER Razil "xiliangMa@outlook.com"
EXPOSE 8080
RUN mkdir -p /usr/share/restapi
COPY ./bin/restapi /usr/share/restapi/
COPY conf /usr/share/restapi/
COPY swagger /usr/share/restapi/
# RUN ln -s /usr/share/restapi/restapi /usr/local/bin/restapi
CMD ["/bin/bash"]