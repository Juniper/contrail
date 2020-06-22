ARG BASE_IMAGE_REGISTRY="opencontrailnightly"
ARG BASE_IMAGE_REPOSITORY="contrail-base"
ARG BASE_IMAGE_TAG="latest"
FROM ${BASE_IMAGE_REGISTRY}/${BASE_IMAGE_REPOSITORY}:${BASE_IMAGE_TAG}

ARG GOPATH
MAINTAINER Nachi Ueno nueno@juniper.net

COPY pip.conf /etc/pip.conf
COPY requirements.txt /requirements.txt
COPY ./src/ $GOPATH/src/github.com/Juniper/
ADD ./contrail /bin/contrailgo
ADD ./contrailcli /bin/contrailcli
ADD ./contrailutil /bin/contrailutil
ADD ./etc /etc/contrail
ADD ./etc/gen_init_psql.sql /usr/share/contrail/init_psql.sql
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

# Upgrade pip used by ansible-deployer
RUN pip install --upgrade pip==19.1.1

RUN pip install -r /requirements.txt

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrailgo", "-c", "/etc/contrail/contrail.yml", "run"]
