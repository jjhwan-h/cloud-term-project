#!/bin/bash

ENV_PATH=${1:-/root/.dev.env}
PRIVATE_KEY_PATH=${2:-/root/cloud-test.pem}

docker run -it\
  -v $LOCAL_PRIVATE_KEY_PATH:$PRIVATE_KEY_PATH \
  -v $LOCAL_ENV_PATH:$ENV_PATH \
  aws