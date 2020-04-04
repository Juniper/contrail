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

RUN yum install -y \
    postgresql \
    git \
    yum clean all -y && \
    rm -rf /var/cache/yum

EXPOSE 9091
WORKDIR /etc/contrail
ENTRYPOINT ["/bin/contrailgo", "-c", "/etc/contrail/contrail.yml", "run"]
