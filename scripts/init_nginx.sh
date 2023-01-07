#!/bin/bash

# Starts nginx in a docker container, given the path to an nginx.conf
# Note the port config of 8080-8081:80-81

NGINX_MAIN_APP="nginx-main"
NETWORK="nginx-app-network"
ROOT_PATH=$(pwd)

if [ "$#" -ne 1 ] || [ ! -f "$ROOT_PATH/$1/nginx.conf" ]; then
  echo "Error: please enter the path to an nginx.conf"
  exit 1
fi

# Create network if not exists
if [[ ! $(docker network ls | grep $NETWORK) ]]; then
    docker network create $NETWORK
fi

if [[ $(docker container ps | grep $NGINX_MAIN_APP) ]]; then
  echo "Container $NGINX_MAIN_APP is running, using $1/nginx.conf."
  bash $ROOT_PATH/scripts/reload_nginx.sh $1
elif [[ ! $(docker container ls -a | grep $NGINX_MAIN_APP) ]]; then
  echo "Creating new container $NGINX_MAIN_APP"
  docker run \
    --name $NGINX_MAIN_APP \
    --hostname $NGINX_MAIN_APP \
    --network $NETWORK \
    -v $ROOT_PATH/html:/usr/share/nginx/html \
    -v $ROOT_PATH/$1/nginx.conf:/etc/nginx/nginx.conf \
    -v $ROOT_PATH/ssl_keys:/usr/share/nginx/ssl_keys \
    -p 8080-8081:80-81 \
    -d \
    nginx
else
  echo "Starting stopped container $NGINX_MAIN_APP"
  docker container start $NGINX_MAIN_APP
  sleep 1
  bash $ROOT_PATH/scripts/reload_nginx.sh $1
fi

# Run nginx if not running
