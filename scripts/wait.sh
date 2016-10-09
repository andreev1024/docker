#!/bin/sh
#WAIT_HOST: mysql
#WAIT_PORT: 3306
#WAIT_START_CMD:
#        "gearmand \
#        --log-file=stderr \
#        --verbose=DEBUG \
#        --queue-type=MySQL \
#        --mysql-host=mysql \
#        --mysql-user=root \
#        --mysql-password=root \
#        --mysql-db=_rebecca \
#        --mysql-table=gearman_queue \
#        --job-retries=2"
#

LOOPS="10"

i=0
while ! nc $WAIT_HOST $WAIT_PORT >/dev/null 2>&1 < /dev/null; do
  i=`expr $i + 1`
  if [ $i -ge $LOOPS ]; then
    echo "$(date) - ${WAIT_HOST}:${WAIT_PORT} still not reachable, giving up"
    exit 1
  fi
  echo "$(date) - waiting for ${WAIT_HOST}:${WAIT_PORT}..."
  sleep 1
done

#continue with further steps

#start the daemon
exec $WAIT_START_CMD