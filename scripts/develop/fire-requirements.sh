#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../..
root=$(pwd)


bash $root/scripts/develop/db/post/start.sh --detach
bash $root/scripts/develop/db/relation/start.sh --detach
bash $root/scripts/develop/db/user/start.sh --detach
bash $root/scripts/develop/nats/start.sh --detach
bash $root/scripts/develop/db/feed/start.sh --detach
bash $root/scripts/develop/db/security/start.sh --detach

exit 0