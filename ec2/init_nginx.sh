#!/bin/bash

NGINX_MAIN_APP="nginx-main-s"
NETWORK="nginx-main-s-network"

# Create network if not exists
if [[ ! $(docker network ls | grep $NETWORK) ]]; then
    docker network create $NETWORK
fi

echo "Creating new container $NGINX_MAIN_APP"
  docker run \
    --name $NGINX_MAIN_APP \
    --hostname $NGINX_MAIN_APP \
    --network $NETWORK \
    -v ~/git/wanna-be/ec2/nginx.conf:/etc/nginx/nginx.conf \
    -p 8090-8093:90-93 \
    -d \
    nginx