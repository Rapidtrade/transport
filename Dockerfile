FROM ubuntu:15.10

MAINTAINER kevin@rapidtrade.biz

EXPOSE  8000

RUN apt-get update && mkdir /opt/transport && mkdir /var/transport
ADD transport /opt/transport
ADD transport.service /lib/systemd/system  

WORKDIR /opt/transport
CMD './transport'