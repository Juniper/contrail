FROM alpine 
MAINTAINER Nachi Ueno nueno@juniper.net

ADD ./contrail /bin/contrail
ADD ./etc /etc/contrail
ADD ./etc/init_mysql.sql /usr/share/contrail/
ADD ./etc/init_psql.sql /usr/share/contrail/
ADD ./public /usr/share/contrail/public
ADD ./templates /usr/share/contrail/templates

RUN apk update && apk add \
        bash \
        mysql-client \
        postgresql-client \
        git \
        python \
        linux-headers \
        py-pip \
        python-dev \
        build-base \
        ansible \
        && pip install --upgrade pip

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c","/etc/contrail/server.yml", "-a", "/etc/contrail/agent.yml", "server"]
