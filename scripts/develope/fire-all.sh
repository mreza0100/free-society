#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../..
root=$(pwd)


bash $root/scripts/develope/db/post/start.sh --detach 
bash $root/scripts/develope/db/relation/start.sh --detach 
bash $root/scripts/develope/db/user/start.sh --detach 
bash $root/scripts/develope/nats/start.sh --detach 

exit 0