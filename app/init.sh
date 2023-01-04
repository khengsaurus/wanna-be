#!/bin/bash

APP="nginx-app"

# Build app image if not exists
if [[ "$(docker images -q $APP:latest 2> /dev/null)" == "" ]]; 
  then docker build app -t $APP
fi

docker run --name go-app -p 8080:8080 -e APP_ID=nginx-app-1 -d $APP