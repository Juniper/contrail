FROM centos:7.4.1708
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

RUN pip install ansible==2.7.11
RUN python2 -m pip install lxml saxpy gevent

# logic to install vnc api client needed by multi-cloud
RUN cd /tmp/contrail-api-client && git fetch --unshallow
RUN cd /tmp/contrail-api-client && scons
RUN cd /tmp/contrail-api-client/build/debug/api-lib && python setup.py install

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c", "/etc/contrail/contrail.yml", "run"]
