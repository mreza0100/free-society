#!/bin/bash 
clear
root=$GOPATH/src/microServiceBoilerplate

cd $root/services/hellgate

gqlgen generate 

echo "Done!"