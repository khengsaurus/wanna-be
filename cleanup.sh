#!/bin/bash

APP="nginx-app"
MAIN="nginx-main"
NETWORK="nginx-app-network"

printf "Stopping containers:\n"

docker container ls -a | grep $APP | cut -c1-14 | xargs docker stop

docker stop $MAIN

sleep 2

printf "\nRemoving containers:\n"

docker container ls -a | grep $APP | cut -c1-14 | xargs docker container rm

docker rm $MAIN

printf "\nRemoving docker network:\n"

docker network rm $NETWORK