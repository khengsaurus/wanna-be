#!/bin/bash

# Takes an argument 1-9 and runs that number of containers, built from ./app/Dockerfile

if [ "$#" -ne 1 ] || ! [[ "$1" =~ ^[1-9]+$ ]]; then
  echo "Error: please enter a number in the range of 1-9"
  exit 1
fi

# -------------------------

APP="expenses-app"
NETWORK="nginx-app-network"

# Build app image if not exists
if [[ "$(docker images -q $APP:latest 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP
fi

# Create network if not exists
if [[ ! $(docker network ls | grep $NETWORK) ]]; then
    docker network create $NETWORK
fi

# Start instances in detached mode, in network
function _run(){
  APP_NAME=$APP-$1
  if [[ $(docker container ps | grep $APP_NAME) ]]; then
    return
  elif [[ $(docker container ls -a | grep $APP_NAME) ]]; then
    echo "Starting stopped container $APP_NAME"
    docker container start $APP_NAME
  else
    echo "Creating new container $APP_NAME"
    docker run \
      --name $APP_NAME \
      --hostname $APP_NAME \
      --network $NETWORK \
      -e APP_ID=$APP_NAME \
      -d \
      $APP
  fi
}

x=1
while [ "$x" -le $1 ]; do
  _run $x
  x=$((x+1))
done