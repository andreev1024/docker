FROM andreev/php5-fpm

MAINTAINER andreev <andreev1024@gmail.com>

RUN apt-get update && \
    apt-get -y install \
        git \
        --no-install-recommends

RUN curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin --filename=composer
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && mkdir -p /var/docker