#
# Copyright (C) 2020 IBM Corporation.
#
# Authors:
# Frederico Araujo <frederico.araujo@ibm.com>
# Teryl Taylor <terylt@ibm.com>

#-----------------------
# Stage: base
#-----------------------
FROM registry.access.redhat.com/ubi8/ubi:8.1-406 as base

ENV PATH=$PATH:/usr/local/go/bin/

ENV GOPATH=/go/

ENV SRC_ROOT=/go/src/github.ibm.com/sysflow/sf-processor/

RUN dnf update -y --disableplugin=subscription-manager && \
     dnf install -y  --disableplugin=subscription-manager wget gcc make git device-mapper-devel

RUN wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz && mkdir -p $SRC_ROOT

COPY core ${SRC_ROOT}core
COPY modules ${SRC_ROOT}/modules
COPY driver ${SRC_ROOT}driver
COPY plugins ${SRC_ROOT}plugins
COPY resources ${SRC_ROOT}resources
COPY Makefile ${SRC_ROOT}

RUN cd ${SRC_ROOT} && make install

#-----------------------
# Stage: runtime
#-----------------------
FROM registry.access.redhat.com/ubi8/ubi-minimal:8.2-267 AS runtime

arg filePath=/sock/sysflow.sock
ENV FILE_PATH=$filepath

arg inputtype=socket
ENV INPUT_TYPE=$inputtype

arg loglevel=info
ENV LOG_LEVEL=$loglevel

ARG VERSION=dev
ARG RELEASE=dev

# Image labels
LABEL "name"="Sysflow Processor"
LABEL "vendor"="IBM"
LABEL "version"="${VERSION}"
LABEL "release"="${RELEASE}"
LABEL "summary"="Sysflow Processor contains a policy engine applying rules on consumed SysFlow stream, and implements an pluggable analytics pipeline"
LABEL "description"="Sysflow Processor contains a policy engine applying rules on consumed SysFlow stream, and implements an pluggable analytics pipeline"
LABEL "io.k8s.display-name"="Sysflow Processor"
LABEL "io.k8s.description"="Sysflow Processor contains a policy engine applying rules on consumed SysFlow stream, and implements an pluggable analytics pipeline"

# Update license
COPY ./LICENSE.md /licenses/

# Copy files from previous stage
COPY --from=base /usr/local/sf-processor/ /usr/local/sf-processor/

# Entrypoint
CMD /usr/local/sf-processor/bin/sfprocessor \
                            ${LOG_LEVEL:+-log} ${LOG_LEVEL} \
                            ${INPUT_TYPE:+-input} ${INPUT_TYPE} \
                            ${FILE_PATH}
