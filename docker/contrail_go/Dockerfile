FROM centos:7.4.1708
MAINTAINER Nachi Ueno nueno@juniper.net

COPY ./contrail /go/src/github.com/Juniper/contrail
ADD ./contrail /bin/contrail
ADD ./contrailcli /bin/contrailcli
ADD ./contrailutil /bin/contrailutil
ADD ./etc /etc/contrail
ADD ./etc/init_mysql.sql /usr/share/contrail/
ADD ./etc/init_psql.sql /usr/share/contrail/
ADD ./etc/init_data.yaml /usr/share/contrail/
ADD ./public /usr/share/contrail/public
ADD ./templates /usr/share/contrail/templates
COPY ./contrail-ansible-deployer /usr/share/contrail/contrail-ansible-deployer

RUN yum update -y &&  yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm  &&\
        yum install -y \
        mysql-client \
        postgresql-client \
        git \
        docker \
        python-requests \
        python-pip

RUN yum remove -y PyYAML

RUN pip install ansible==2.4.2.0

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c", "/etc/contrail/contrail.yml", "run"]
