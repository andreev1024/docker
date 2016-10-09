#!/usr/bin/env bash

NAME="deploy"
ID="1000"

groupadd -g $ID $NAME && useradd -m -g $NAME -s /bin/bash -u $ID $NAME

exec /sbin/my_init