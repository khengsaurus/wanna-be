#!/bin/bash

APP="nhttpd"
NETWORK="nhttpd"

printf "Stopping containers:\n"

docker container ls -a | grep $APP | cut -c1-12 | xargs docker stop

sleep 1

printf "\nRemoving containers:\n"

docker container ls -a | grep $APP | cut -c1-12 | xargs docker container rm

printf "\nRemoving docker networks:\n"

docker network ls | grep $NETWORK | cut -c1-12 | xargs docker network rm
