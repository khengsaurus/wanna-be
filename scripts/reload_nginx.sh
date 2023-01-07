#!/bin/bash

NGINX_MAIN_APP="nginx-main"

if [[ ! -f "$(pwd)/$1/nginx.conf" ]]; then
  echo "Please enter a dir in $(pwd) with an nginx.conf"
  exit 1
fi

docker cp $(pwd)/$1/nginx.conf $NGINX_MAIN_APP:/etc/nginx/nginx.conf

docker exec $NGINX_MAIN_APP nginx -s reload