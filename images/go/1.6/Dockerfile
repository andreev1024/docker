FROM phusion/baseimage:0.9.18

MAINTAINER andreev <andreev1024@gmail.com>

RUN apt-get update && \
    apt-get -y install \
        git \
        --no-install-recommends && \
        apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz && \
    tar -C /usr/local/ -xzf go1.6.linux-amd64.tar.gz && \
    rm go1.6.linux-amd64.tar.gz

RUN cd / && mkdir go go/src go/pkg go/bin

WORKDIR /go

ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH /go

CMD ["/sbin/my_init"]