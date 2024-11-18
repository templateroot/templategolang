FROM scratch
WORKDIR /app
COPY ./wjgotemplate1 ./
COPY ./WJGOTEMPLATE1.cfg ./
ADD ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/app/wjgotemplate1"]