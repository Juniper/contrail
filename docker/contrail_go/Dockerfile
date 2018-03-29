FROM alpine 
MAINTAINER Nachi Ueno nueno@juniper.net

ADD ./contrail /bin/contrail
ADD ./etc /etc/contrail
ADD ./etc/init_mysql.sql /usr/share/contrail/
ADD ./etc/init_psql.sql /usr/share/contrail/

RUN apk update && apk add mysql-client postgresql-client

EXPOSE 9091
ENTRYPOINT ["/bin/contrail","-c","/etc/contrail/server.yml","server"]
