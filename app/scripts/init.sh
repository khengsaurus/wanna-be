#!/bin/bash

APP="nginx-go-app"

# Build app image if not exists
if [[ "$(docker images -q $APP:latest 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP
fi

docker run --name $APP-s -p 8080:8080 -e APP_ID=$APP-s -d $APP