#
# Copyright (C) 2021 IBM Corporation.
#
# Authors:
# Frederico Araujo <frederico.araujo@ibm.com>
# Teryl Taylor <terylt@ibm.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#-----------------------
# Stage: base
#-----------------------
ARG UBI_VER=8.4-206
FROM registry.access.redhat.com/ubi8/ubi:${UBI_VER} AS base

# Environment and build args
ARG VERSION=dev

ARG BUILD_NUMBER=dev

ENV PATH=$PATH:/usr/local/go/bin/

ENV GOPATH=/go/

ENV SRC_ROOT=/go/src/github.com/sysflow-telemetry/sf-processor/

# Install dependencies
RUN dnf update -y --disableplugin=subscription-manager && \
     dnf install -y  --disableplugin=subscription-manager wget gcc make git device-mapper-devel

RUN wget https://dl.google.com/go/go1.16.3.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz && mkdir -p $SRC_ROOT

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
FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4-200 AS runtime

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
LABEL "maintainer"="The SysFlow team"
LABEL "vendor"="SysFlow"
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
RUN microdnf -y update && \
    ( microdnf -y clean all ; rm -rf /var/cache/{dnf,yum} ; true ) && \
    mkdir -p /sock && chown -R 1001:1001 /sock
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
