#!/bin/bash 
clear




DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR
cd ../..
ROOT=$(pwd)

cd $ROOT

cd ./services/hellgate

gqlgen generate 

# echo "Done!"