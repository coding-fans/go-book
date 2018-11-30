FROM scratch
ENV TZ=Asia/Shanghai
ADD zoneinfo.tar.gz /
ADD cacert.pem /etc/ssl/certs/
ADD demo /
CMD ["/demo"]
