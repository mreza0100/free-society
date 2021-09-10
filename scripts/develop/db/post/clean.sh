#!/bin/bash


DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR
cd ../../../../
ROOT=$(pwd)


docker container exec -ti docker_post_service_mongo_1 bash -c "printf 'use posts \n db.dropDatabase()' | mongo"

