#!/bin/bash 
clear
root=$GOPATH/src/microServiceBoilerplate


cd $root/services/user/docker
docker-compose up
