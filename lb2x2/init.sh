#!/bin/bash

# Start 2x nginx balancing 2x app each

APP="nginx-app"
MAIN="nginx-main"
NETWORK="nginx-app-network"

# Build app image if not exists
if [[ "$(docker images -q $APP:latest 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP
fi

# Create network
docker network create $NETWORK

# Start 3 instances in detached mode
function _run(){
  docker run \
  --name $APP-$1 \
  --hostname $APP-$1 \
  --network $NETWORK \
  -d \
  $APP
}

_run 0
_run 1
_run 2
_run 3

# Run nginx
docker run \
  --name $MAIN \
  --hostname $MAIN \
  --network $NETWORK \
  -p 8080-8081:8080-8081 \
  -v /Users/kheng/git/wanna-be/lb2x2/nginx.conf:/etc/nginx/nginx.conf \
  -d \
  nginx
