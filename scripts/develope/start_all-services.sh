#!/bin/bash 
clear
root=$GOPATH/src/microServiceBoilerplate


bash $root/scripts/develope/services/start.hellgate.sh &
bash $root/scripts/develope/services/start.user.sh
