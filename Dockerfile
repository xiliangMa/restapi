FROM 47.244.62.249:9090/maxl/golang:1.11

MAINTAINER xiliangMa "xiliangMa@outlook.com"

EXPOSE 8080
EXPOSE 8088
EXPOSE 10443

WORKDIR /usr/share/restapi
RUN mkdir -p /usr/share/restapi/conf && \
    mkdir -p /usr/share/restapi/swagger

COPY ./bin/restapi /usr/share/restapi/
COPY conf /usr/share/restapi/conf
COPY swagger /usr/share/restapi/swagger

# RUN ln -s /usr/share/restapi/restapi /usr/local/bin/restapi
CMD ["./restapi"]
