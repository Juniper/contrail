ARG REGISTRY="opencontrailnigtly"
ARG REPOSITORY="contrail-base"
ARG TAG="latest"
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

# creating link as needed by multi-cloud
RUN ln -s /usr/share/contrail/contrail-ansible-deployer /tmp/

# apisrv pkg dependencies
RUN yum install -y mysql postgresql

# deploy pkg dependencies
RUN yum install -y git

# contrail-ansible-deployer dependencies
RUN yum install -y python-requests
RUN easy_install pip
RUN pip install ansible==2.7.11

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrail", "-c", "/etc/contrail/contrail.yml", "run"]
