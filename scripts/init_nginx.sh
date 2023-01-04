#!/bin/bash

# Starts nginx in a docker container, given the path to an nginx.conf

MAIN="nginx-main"
NETWORK="nginx-app-network"

if [ "$#" -ne 2 ] || [ ! -f "$1" ]; then
  echo "Error: please enter the path to an nginx.conf, followed by the port config"
  exit 1
fi

# Run nginx
docker run \
  --name $MAIN \
  --hostname $MAIN \
  --network $NETWORK \
  -v /$1:/etc/nginx/nginx.conf \
  -p $2 \
  -d \
  nginx
