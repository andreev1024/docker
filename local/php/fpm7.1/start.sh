#!/usr/bin/env bash
php5enmod imap;
php5enmod mcrypt;

groupadd -g $MY_ID $MY_NAME && useradd -m -g $MY_NAME -s /bin/bash -u $MY_ID $MY_NAME
usermod -a -G $MY_NAME www-data
usermod -a -G www-data $MY_NAME

# composer init
#composer self-update
#sudo -H -u $MY_NAME bash -c 'composer global install'

# put hosts to /etc/hosts
HOST_NAMES=($HOST_NAMES)

for i in "${HOST_NAMES[@]}"
do
  echo "$HOST_IP       $i" >> /etc/hosts
done

$CUSTOM_CMD
exec $AFTER_SCRIPT_CMD
