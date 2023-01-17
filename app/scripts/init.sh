#!/bin/bash

APP="expenses-app"
TAG="1"

# Build app image if not exists
if [[ "$(docker images -q $APP:$TAG 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP:$TAG
fi

docker run --name $APP-s -p 8080:8080 -e APP_ID=$APP-s -d $APP:$TAG