#!/bin/bash
clear

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR




echo "Feed"
bash ./feed/clean.sh
echo "----"
sleep 1s

echo "post"
bash ./post/clean.sh
echo "----"
sleep 1s

echo "relation"
bash ./relation/clean.sh
echo "----"
sleep 1s

echo "security"
bash ./security/clean.sh
echo "----"
sleep 1s

echo "user"
bash ./user/clean.sh

echo "notification"
bash ./notification/clean.sh
