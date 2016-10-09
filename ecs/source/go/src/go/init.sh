#!/usr/bin/env bash

# Copy ssh keys
cp -R /var/docker/.ssh /root/.ssh

cd /go/src/app && go install -v && app;