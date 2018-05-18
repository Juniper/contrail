FROM centos:7.4.1708
MAINTAINER Nachi Ueno nueno@juniper.net

ADD ./contrail /bin/contrail
ADD ./contrailcli /bin/contrailcli
ADD ./etc /etc/contrail
ADD ./etc/init_mysql.sql /usr/share/contrail/
ADD ./etc/init_psql.sql /usr/share/contrail/
ADD ./public /usr/share/contrail/public
ADD ./templates /usr/share/contrail/templates

RUN yum update -y &&  yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm  &&\
        yum install -y \
        mysql-client \
        postgresql-client \
        git \
        ansible \
        docker \
        python-pip \
        && pip install --upgrade pip

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c", "/etc/contrail/contrail.yml", "run"]
