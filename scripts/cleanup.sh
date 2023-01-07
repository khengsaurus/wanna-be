#!/bin/bash

APP="nginx-go-app"
NGINX_MAIN_APP="nginx-main"
NETWORK="nginx-app-network"

printf "Stopping containers:\n"

docker container ls -a | grep $APP | cut -c1-12 | xargs docker stop

docker stop $NGINX_MAIN_APP

sleep 1

printf "\nRemoving containers:\n"

docker container ls -a | grep $APP | cut -c1-12 | xargs docker container rm

docker rm $NGINX_MAIN_APP

printf "\nRemoving docker network:\n"

docker network rm $NETWORK