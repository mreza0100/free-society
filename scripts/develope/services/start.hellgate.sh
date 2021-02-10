#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../../..
root=$(pwd)





export GIN_MODE=debug
export MODE=dev
export SECRET_KEY="ap:OUWE#@#9iwjd@u3wj20i2erakwjdfAOJGF_@!I"

go run $root/services/hellgate/server
