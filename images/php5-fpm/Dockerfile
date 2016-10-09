FROM phusion/baseimage:0.9.17

MAINTAINER andreev <andreev1024@gmail.com>

RUN apt-get update && \
    apt-get -y install \
        php5 \
        php5-dev \
        php5-cli \
        php5-fpm \
        php5-json \
        php5-sqlite \
        php5-curl \
        php5-gd \
        php5-mcrypt \
        php5-intl \
        php5-imap \
        php5-gearman \
        php5-memcache \
        php5-mysqlnd \
        php5-readline \
        nano \
        --no-install-recommends

RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

CMD ["/sbin/my_init", "--", "php5-fpm", "--nodaemonize"]