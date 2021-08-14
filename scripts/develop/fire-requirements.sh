#!/bin/bash
clear

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR
cd ../..
ROOT=$(pwd)


bash $ROOT/scripts/develop/db/post/start.sh --detach
bash $ROOT/scripts/develop/db/relation/start.sh --detach
bash $ROOT/scripts/develop/db/user/start.sh --detach
bash $ROOT/scripts/develop/nats/start.sh --detach
bash $ROOT/scripts/develop/db/feed/start.sh --detach
bash $ROOT/scripts/develop/db/security/start.sh --detach

exit 0