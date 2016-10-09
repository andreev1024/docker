FROM phusion/baseimage:0.9.17

MAINTAINER andreev <andreev1024@gmail.com>

CMD ["/sbin/my_init"]

RUN curl -sL https://deb.nodesource.com/setup_0.12 | sudo bash -
RUN apt-get update && apt-get -y install \
    make \
    g++ \
    nodejs \
    python \
    nano \
    git \
    --no-install-recommends

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*