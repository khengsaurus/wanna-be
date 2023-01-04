#!/bin/bash

if [ "$#" -ne 2 ] || [ ! -f "$2" ]; then
  echo "Error: please enter the container name/id, folllowed by the path to an nginx.conf"
  exit 1
fi

docker cp $2 $1:/etc/nginx/nginx.conf

docker exec $1 nginx -s reload