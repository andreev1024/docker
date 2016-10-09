FROM golang:1.6.0

MAINTAINER andreev <andreev1024@gmail.com>

RUN apt-get update && \
    apt-get -y install \
        python-pip \
        python-dev \
        --no-install-recommends && \
        apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENV TERM xterm
ENV SHELL /bin/bash

RUN pip install credstash awscli