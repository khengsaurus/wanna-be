#!/bin/bash

# Takes an argument 1-9 and runs that number of containers, built from ./app/Dockerfile

if [ "$#" -ne 1 ] || ! [[ "$1" =~ ^[1-9]+$ ]]; then
  echo "Error: please enter a number in the range of 1-9"
  exit 1
fi

# -------------------------

APP="nginx-app"
MAIN="nginx-main"
NETWORK="nginx-app-network"

# Build app image if not exists
if [[ "$(docker images -q $APP:latest 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP
fi

# Create network
docker network create $NETWORK

# Start instances in detached mode, in network
function _run(){
  docker run \
    --name $APP-$1 \
    --hostname $APP-$1 \
    --network $NETWORK \
    -e APP_ID=$APP-$1 \
    -d \
    $APP
}

x=1
while [ "$x" -le $1 ]; do
  _run $x
  x=$((x+1))
done