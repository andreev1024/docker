FROM phusion/baseimage:0.9.19

MAINTAINER andreev <andreev1024@gmail.com>

RUN add-apt-repository ppa:ondrej/php
RUN apt-get update && \
    apt-get -y --allow-unauthenticated install \
        php7.1-common \
        php7.1-cli \
        php7.1-dev \
        php7.1-fpm \
        php7.1-curl \
        php7.1-sqlite3 \
        php7.1-json \
        php7.1-tidy \
        php7.1-mysql \
        php7.1-gd \
        php7.1-mcrypt \
        php7.1-intl \
        php7.1-imap \
        php7.1-readline \
        nano \
        --no-install-recommends

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

CMD ["/sbin/my_init", "--", "php7.1-fpm", "--nodaemonize"]