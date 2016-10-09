FROM phusion/baseimage:0.9.15

MAINTAINER andreev <andreev1024@gmail.com>

EXPOSE 4730

RUN apt-get update -qq && \
    apt-get install -y gcc make g++ libboost-all-dev gperf libevent-dev \
        uuid-dev libmysqld-dev libpq-dev libsqlite3-dev libhiredis-dev && \
    mkdir /tmp/gearman && \
    curl -L --insecure --silent https://launchpad.net/gearmand/1.2/1.1.12/+download/gearmand-1.1.12.tar.gz \
        | tar -xz --strip 1 -C /tmp/gearman && \
    cd /tmp/gearman && \
    ./configure \
        --prefix=/usr \
        --sysconfdir=/etc \
        --localstatedir=/var \
        --enable-hiredis && \
    make && \
    make install && \
    adduser --system --no-create-home --disabled-login --disabled-password --group gearman && \
    apt-get purge -y gcc make gperf libevent-dev uuid-dev \
        libmysqld-dev libpq-dev libsqlite3-dev libhiredis-dev  && \
    apt-get autoremove -y && \
    apt-get install -y telnet libmysqlclient18 libhiredis0.10 libevent-2.0-5 libpq5 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ENTRYPOINT ["/sbin/my_init"]
CMD ["--", "gearmand", "--log-file=stderr"]