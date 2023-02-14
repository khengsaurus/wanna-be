#!/bin/bash

# 1 - path to Dockerfile
# 2 - image name
function build_img_if_not_exists(){
  if [[ "$(docker images -q $2 2> /dev/null)" == "" ]]; 
    then docker build $1 -t $2
  fi
}

# Mahjong CMS
MJ_CMS_REPO=~/git/mahjong-cms
MJ_CMS_NAME=mj-cms
MJ_CMS_IMG=mj-cms:1

build_img_if_not_exists $MJ_CMS_REPO $MJ_CMS_IMG
docker build $MJ_CMS_REPO -t $MJ_CMS_IMG
docker run \
  --name $MJ_CMS_NAME \
  --hostname $MJ_CMS_NAME \
  -e APP_ID=$MJ_CMS_NAME-main \
  -p 8090:8080 \
  -d \
  $MJ_CMS_IMG

# Expenses app
EX_APP_REPO=~/git/wanna-be/app
EX_APP_NAME=expenses-app
EX_APP_IMG=expenses-app:1

build_img_if_not_exists $EX_APP_REPO $EX_APP_IMG
docker build $EX_APP_REPO -t $EX_APP_IMG
docker run \
  --name $EX_APP_NAME \
  --hostname $EX_APP_NAME \
  -e APP_ID=$EX_APP_NAME-main \
  -p 8092:8080 \
  -d \
  $EX_APP_IMG

