ARG BASE_IMAGE_REGISTRY="opencontrailnightly"
ARG BASE_IMAGE_REPOSITORY="contrail-base"
ARG BASE_IMAGE_TAG="latest"
FROM ${BASE_IMAGE_REGISTRY}/${BASE_IMAGE_REPOSITORY}:${BASE_IMAGE_TAG}

ARG GOPATH
MAINTAINER Nachi Ueno nueno@juniper.net

COPY pip.conf /etc/pip.conf
COPY ./src/ $GOPATH/src/github.com/Juniper/
ADD ./contrail /bin/contrailgo
ADD ./contrailcli /bin/contrailcli
ADD ./contrailutil /bin/contrailutil
ADD ./etc /etc/contrail
ADD ./etc/gen_init_psql.sql /usr/share/contrail/gen_init_psql.sql
ADD ./etc/init_psql.sql /usr/share/contrail/init_psql.sql
ADD ./etc/init_data.yaml /usr/share/contrail/
ADD ./public /usr/share/contrail/public
ADD ./templates /usr/share/contrail/templates
COPY ./contrail-ansible-deployer /usr/share/contrail/contrail-ansible-deployer

# creating link as needed by multi-cloud
RUN ln -s /usr/share/contrail/contrail-ansible-deployer /tmp/

RUN yum install -y \
    postgresql \
    git \
    python-requests python-pip && \
    yum clean all -y

RUN pip install MarkupSafe==1.1.1 PyYAML==5.3 ansible==2.7.11 bcrypt==3.1.7 cffi==1.14.0 cryptography==2.8 enum34==1.1.6 jinja2==2.11.1 paramiko==2.7.1 pycparser==2.19 pynacl==1.3.0

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrailgo", "-c", "/etc/contrail/contrail.yml", "run"]
