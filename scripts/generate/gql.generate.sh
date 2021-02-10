#!/bin/bash 
clear




dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../..
root=$(pwd)

cd $root

cd ./services/hellgate

gqlgen generate 

# echo "Done!"