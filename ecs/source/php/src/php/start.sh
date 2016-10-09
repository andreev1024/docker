#!/usr/bin/env bash
php5enmod imap;
php5enmod mcrypt;

# create user/groups
DEPLOY_ID=2000

groupadd -g $DEPLOY_ID deploy && useradd -m -g deploy -s /bin/bash -u $DEPLOY_ID deploy

usermod -a -G deploy www-data

# composer init
composer self-update
sudo -H -u composer bash -c 'composer global install'

eval $CUSTOM_CMD

exec $AFTER_SCRIPT_CMD