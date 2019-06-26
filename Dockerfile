FROM golang:1.11

MAINTAINER xiliangMa "xiliangMa@outlook.com"

EXPOSE 8080
EXPOSE 8088
EXPOSE 10443

ENV MARIADB_HOST localhost
ENV MARIADB_USER restapi
ENV MARIADB_PASSWORD restapi

WORKDIR /usr/share/restapi
RUN mkdir -p /usr/share/restapi/conf && \
    mkdir -p /usr/share/restapi/swagger


COPY entrypoint.sh /usr/share/restapi/
COPY ./bin/restapi /usr/share/restapi/
COPY conf /usr/share/restapi/conf
COPY swagger /usr/share/restapi/swagger

RUN chmod +x /usr/share/restapi/entrypoint.sh


# RUN ln -s /usr/share/restapi/restapi /usr/local/bin/restapi
ENTRYPOINT ["sh", "/usr/share/restapi/entrypoint.sh" ]
