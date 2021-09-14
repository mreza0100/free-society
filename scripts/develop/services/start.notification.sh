#!/bin/bash
clear

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../../
ROOT=$(pwd)


if [[ "$1" == "--from-modd" ]];
then
      source $ROOT/scripts/develop/env.sh
      if ! go run $ROOT/services/notification/server/server.go
      then
            # there was an error from go program
            tput bel
            sleep 5s
      fi
fi

if [[ "$1" == "" ]];
then
      cd $ROOT/services/notification && modd -bn
fi