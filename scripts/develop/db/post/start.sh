#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../../../
ROOT=$(pwd)


cd $ROOT/services/post/docker

sudo docker-compose up $1

