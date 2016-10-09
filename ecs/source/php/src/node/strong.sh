#!/usr/bin/env bash
NODE_APP_PATH=($NODE_APP_PATH)
NODE_APPS=($NODE_APPS)

for node_one_app in "${NODE_APPS[@]}"
do
    cd $NODE_APP_PATH/$node_one_app && slc build --install && slc start
done

exec $AFTER_SCRIPT_CMD