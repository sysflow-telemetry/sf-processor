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

# Environment and build args
ARG VERSION=dev

ARG BUILD_NUMBER=dev

ENV PATH=$PATH:/usr/local/go/bin/

ENV GOPATH=/go/

ENV SRC_ROOT=/go/src/github.com/sysflow-telemetry/sf-processor/

# Install dependencies
RUN dnf update -y --disableplugin=subscription-manager && \
     dnf install -y  --disableplugin=subscription-manager wget gcc make git device-mapper-devel

RUN wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.14.2.linux-amd64.tar.gz && mkdir -p $SRC_ROOT

# Copy sources
COPY core ${SRC_ROOT}core
COPY driver ${SRC_ROOT}driver
COPY plugins ${SRC_ROOT}plugins
COPY resources ${SRC_ROOT}resources
COPY Makefile ${SRC_ROOT}
COPY makefile.manifest.inc ${SRC_ROOT}

# Build
RUN cd ${SRC_ROOT} && \
    make SYSFLOW_VERSION=$VERSION \
         SYSFLOW_BUILD_NUMBER=$BUILD_NUMBER \
         install

#-----------------------
# Stage: runtime
#-----------------------
FROM registry.access.redhat.com/ubi8/ubi-minimal:8.2-267 AS runtime

# Environment and build args
ARG VERSION=dev

ARG RELEASE=dev

ARG inputpath=/sock/sysflow.sock
ENV INPUT_PATH=$inputpath

ARG driver=socket
ENV DRIVER=$driver

ARG driver_dir=/usr/local/sysflow/resources/drivers
ENV DRIVER_DIR=$driver_dir

ARG plugin_dir=/usr/local/sysflow/resources/plugins
ENV PLUGIN_DIR=$plugin_dir

ARG loglevel=info
ENV LOG_LEVEL=$loglevel

ARG configpath=/usr/local/sysflow/conf/pipeline.json
ENV CONFIG_PATH=$configpath

# Image labels
LABEL "name"="SysFlow Processor"
LABEL "vendor"="IBM"
LABEL "version"="${VERSION}"
LABEL "release"="${RELEASE}"
LABEL "summary"="SysFlow Processor implements a pluggable stream-processing pipeline and contains a built-in policy engine that evaluates rules on the ingested SysFlow stream"
LABEL "description"="SysFlow Processor implements a pluggable stream-processing pipeline and contains a built-in policy engine that evaluates rules on the ingested SysFlow stream"
LABEL "io.k8s.display-name"="SysFlow Processor"
LABEL "io.k8s.description"="SysFlow Processor implements a pluggable stream-processing pipeline and contains a built-in policy engine that evaluates rules on the ingested SysFlow stream"

# Update license
COPY ./LICENSE.md /licenses/

# Copy files from previous stage
COPY --from=base --chown=1001:1001 /usr/local/sysflow/ /usr/local/sysflow/
RUN mkdir -p /sock && chown -R 1001:1001 /sock
VOLUME /sock
USER 1001

# Entrypoint
CMD /usr/local/sysflow/bin/sfprocessor \
                            ${LOG_LEVEL:+-log} ${LOG_LEVEL} \
                            ${DRIVER:+-driver} ${DRIVER} \
                            ${DRIVER_DIR:+-driverdir} ${DRIVER_DIR} \
                            ${PLUGIN_DIR:+-plugdir} ${PLUGIN_DIR} \
                            ${CONFIG_PATH:+-config} ${CONFIG_PATH} \
                            ${INPUT_PATH}
