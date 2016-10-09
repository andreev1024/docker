#!/usr/bin/env bash

#   Run migrate sript for https://github.com/mattes/migrate.
#   It parse .env file and run migrate.
#
#   .env file example:
#
#       ...
#       APP_PORT=8080
#       APP_DOMAIN=btcplate.net
#       APP_SCHEMA=http
#       ...
#
#   Input args
#      1)   path to env file;
#      2)   path to migrations;
#

VARIABLES=("DB_USER DB_PASSWORD DB_HOST DB_PORT DB_NAME DB_DRIVER")
#/go/src/app/migrations
PATH_TO_ENV_FILE=$1
PATH_TO_MIGRATIONS=$2

for VARIABLE_NAME in $VARIABLES
do
    ROW=$(cat $PATH_TO_ENV_FILE|grep $VARIABLE_NAME)
    ROW_ARRAY=(${ROW//=/ })
    declare $VARIABLE_NAME="${ROW_ARRAY[1]}"
    if [ -z ${!VARIABLE_NAME} ]; then
        echo "env variable ($VARIABLE_NAME) is undefined"
        exit 1
    fi
done

migrate -url "$DB_DRIVER://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" -path $PATH_TO_MIGRATIONS up