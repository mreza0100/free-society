#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir





bash ./feed/clean.sh
sleep 1s

bash ./post/clean.sh
sleep 1s

bash ./relation/clean.sh
sleep 1s

bash ./security/clean.sh
sleep 1s

bash ./user/clean.sh
