#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../../..
root=$(pwd)


source $root/scripts/develope/env.sh


go run $root/services/hellgate/server
