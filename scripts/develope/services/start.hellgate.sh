#!/bin/bash 
root=$GOPATH/src/microServiceBoilerplate

clear

go run $root/services/hellgate/server
