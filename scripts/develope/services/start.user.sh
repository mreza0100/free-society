#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../../..
root=$(pwd)


export SALT="saltsalt___salt"
export MODE=dev


go run $root/services/user/server
