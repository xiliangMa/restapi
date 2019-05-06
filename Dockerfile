FROM golang:1.11
EXPOSE 80
COPY ./bin/restapi /usr/local/bin/
CMD ["restapi"]