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
COPY ./contrail-api-client /tmp/contrail-api-client

# creating link as needed by multi-cloud
RUN ln -s /usr/share/contrail/contrail-ansible-deployer /tmp/

RUN yum update -y &&  yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm  &&\
        yum install -y \
        mysql-client \
        postgresql-client \
        git \
        docker \
        python-requests \
        python-pip \
        gcc \
        scons

RUN yum remove -y PyYAML

RUN pip install ansible==2.4.2.0
RUN python2 -m pip install lxml saxpy gevent

# logic to install vnc api client needed by multi-cloud
RUN cd /tmp/contrail-api-client && git fetch --unshallow
RUN cd /tmp/contrail-api-client && scons
RUN cd /tmp/contrail-api-client/build/debug/api-lib && python setup.py install

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c", "/etc/contrail/contrail.yml", "run"]
