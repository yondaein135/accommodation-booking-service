#!/bin/bash

#./generate.sh
#if [ $? -ne 0 ]; then
#    echo "generate.sh failed to execute properly"
#    exit 1
#fi

NETWORK_NAME="abs-network"

if [ -z $(docker network ls --filter name=^${NETWORK_NAME}$ --format="{{ .Name }}") ] ; then
  echo "Creating network ${NETWORK_NAME}"
  docker network create ${NETWORK_NAME}
else
  echo "Network ${NETWORK_NAME} already exists"
fi


echo "Starting Database Service..."
docker-compose -f db-docker-compose.yml up -d

docker-compose -f services-docker-compose.yml up --build