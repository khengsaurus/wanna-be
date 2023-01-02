#!/bin/bash

# Start 1xnginx + 1xapp

APP="nginx-app"
MAIN="nginx-main"
NETWORK="nginx-app-network"

# Build app image if not exists
if [[ "$(docker images -q $APP:latest 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP
fi

# Create network
docker network create $NETWORK

docker run \
  --name $APP \
  --hostname $APP \
  --network $NETWORK \
  -d \
  $APP


# Run nginx
docker run \
  --name $MAIN \
  --hostname $MAIN \
  --network $NETWORK \
  -p 8080:80 \
  -v /Users/kheng/git/wanna-be/single/nginx.conf:/etc/nginx/nginx.conf \
  -d \
  nginx
