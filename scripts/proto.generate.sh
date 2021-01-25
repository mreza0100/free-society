#!/bin/bash 
clear
root=$GOPATH/src/vv
generated=$root/proto/generated
raw_protos=$root/proto/raw_protos


function generate {
	protoc --go_out=$generated/$1 --go_opt=paths=source_relative \
		--go-grpc_out=$generated/$1 --go-grpc_opt=paths=source_relative \
		./*.proto
}

protos=("user")


for t in ${protos[@]};
do
	cd $root
	mkdir -p $generated/$t
	cd $raw_protos/$t
	generate $t
	
done




echo "Successfuly generated"