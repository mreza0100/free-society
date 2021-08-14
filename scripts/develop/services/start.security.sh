#!/bin/bash
clear

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../../
ROOT=$(pwd)


source $ROOT/scripts/develop/env.sh


go run $ROOT/services/security/server/server.go
