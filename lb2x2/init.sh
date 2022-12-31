#!/bin/bash

# Start 2x nginx balancing 2x app each

APP="nginx-app"
MAIN="nginx-main"
NETWORK="nginx-app-network"

function _run(){
  docker run \
  --name $APP-$1 \
  --hostname $APP-$1 \
  --network $NETWORK \
  -d \
  $APP
}

# Build image
docker build -t $APP app

# Create network
docker network create $NETWORK

# Start 3 instances in detached mode
_run 1
_run 2
_run 3
_run 4

docker run \
  --name $MAIN \
  --hostname $MAIN \
  --network $NETWORK \
  -p 80-81:8080-8081 \
  -v /Users/kheng/git/wanna-be/lb2x2/nginx.conf:/etc/nginx/nginx.conf \
  -d \
  nginx

# nginx with a html path
# docker run \
#   --name nginx \
#   --hostname ng1 \
#   -p 80:80 \
#   -v html:/usr/share/nginx/html \
#   -d \
#   nginx
