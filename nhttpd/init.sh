#!/bin/bash

IMAGE_NAME="nhttpd"
NETWORK_BE="nhttpd-be"
NETWORK_FE="nhttpd-fe"
NETWORK_SUBNET_BE="10.0.0.0/24"
NETWORK_SUBNET_FE="10.0.1.0/24"

# Build image if not exists
if [[ "$(docker images -q $IMAGE_NAME:latest 2> /dev/null)" == "" ]]; 
  then docker build nhttpd -t $IMAGE_NAME
fi

# Create networks
docker network create $NETWORK_BE --subnet $NETWORK_SUBNET_BE
docker network create $NETWORK_FE --subnet $NETWORK_SUBNET_FE

# Start 2 instances in detached mode
function _run(){
docker run \
  --name $IMAGE_NAME-$1 \
  --hostname $IMAGE_NAME-$1 \
  --network $2 \
  --cap-add=NET_ADMIN \
  -d \
  $IMAGE_NAME
}

_run 0 $NETWORK_BE
_run 1 $NETWORK_FE

# Start gateway
docker run \
  --name $IMAGE_NAME-gw \
  -d \
  $IMAGE_NAME

# Connect gateway to both be and fe networks
docker network connect $NETWORK_BE $IMAGE_NAME-gw
docker network connect $NETWORK_FE $IMAGE_NAME-gw

# Add routes in nhttpd apps and edit content
docker exec nhttpd-0 ip route add 10.0.1.0/24 via 10.0.0.3
docker exec nhttpd-0 bash -c "echo 'Hello from nhttpd-0' > htdocs/index.html"
docker exec nhttpd-1 ip route add 10.0.0.0/24 via 10.0.1.3
docker exec nhttpd-1 bash -c "echo 'Hello from nhttpd-1' > htdocs/index.html"

# Test:
# > docker exec -it nhttpd-0 bash
# > curl curl 10.0.1.2