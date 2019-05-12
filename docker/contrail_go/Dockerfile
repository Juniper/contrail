ARG REGISTRY="docker.io"
ARG REPOSITORY="centos"
ARG TAG="7.4.1708"
FROM ${REGISTRY}/${REPOSITORY}:${TAG}

ARG GOPATH
MAINTAINER Nachi Ueno nueno@juniper.net

COPY ./src/ $GOPATH/src/github.com/Juniper/
ADD ./contrail /bin/contrail
ADD ./contrailcli /bin/contrailcli
ADD ./contrailutil /bin/contrailutil
ADD ./etc /etc/contrail
ADD ./etc/gen_init_mysql.sql /usr/share/contrail/init_mysql.sql
ADD ./etc/gen_init_psql.sql /usr/share/contrail/init_psql.sql
ADD ./etc/init_data.yaml /usr/share/contrail/
ADD ./public /usr/share/contrail/public
ADD ./templates /usr/share/contrail/templates
COPY ./contrail-ansible-deployer /usr/share/contrail/contrail-ansible-deployer
COPY ./contrail-api-client /tmp/contrail-api-client

# creating link as needed by multi-cloud
RUN ln -s /usr/share/contrail/contrail-ansible-deployer /tmp/

RUN yum update -y &&  yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm  &&\
        yum install -y \
        mysql-client \
        postgresql-client \
        git \
        python-requests \
        python-pip \
        gcc \
        scons

RUN yum remove -y PyYAML

# Latest versions of ansible (>2.6) and lower versions (<2.5) have issues with
# adding a new compute to a cluster. Stick with 2.5.2.0 until Ansible fixes the
# issues. Refer: https://bugs.launchpad.net/kolla-ansible/+bug/1784123
RUN pip install ansible==2.5.2.0
RUN python2 -m pip install lxml saxpy gevent

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c", "/etc/contrail/contrail.yml", "run"]
