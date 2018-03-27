FROM scratch
MAINTAINER Nachi Ueno nueno@juniper.net

ADD ./contrail /bin/contrail
ADD ./etc /etc/contrail

EXPOSE 9091
ENTRYPOINT ["/bin/contrail","-c","/etc/contrail/server.yml","server"]
