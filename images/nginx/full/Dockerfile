FROM phusion/baseimage:0.9.18

MAINTAINER andreev <andreev1024@gmail.com>

RUN add-apt-repository ppa:nginx/stable && \
    apt-get update && \
    apt-get -y install nginx-full && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# nginx configuration
RUN echo "daemon off;" >> /etc/nginx/nginx.conf
RUN ln -sf /dev/stdout /var/log/nginx/access.log
RUN ln -sf /dev/stderr /var/log/nginx/error.log

# adding additional daemons
RUN mkdir -p /etc/service/nginx
ADD start.sh /etc/service/nginx/run
RUN chmod +x /etc/service/nginx/run

CMD ["/sbin/my_init"]

EXPOSE 80