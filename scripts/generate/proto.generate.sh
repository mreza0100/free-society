#!/bin/bash 
clear




dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../..
root=$(pwd)

generated=$root/proto/generated
raw_protos=$root/proto/raw_protos


function generate {
	protoc --go_out=$generated/$1 --go_opt=paths=source_relative \
		--go-grpc_out=$generated/$1 --go-grpc_opt=paths=source_relative \
		./*.proto
}

protos=("user" "post" "relation" "feed" "security" "nats")


for t in ${protos[@]};
do
	cd $root
	mkdir -p $generated/$t
	cd $raw_protos/$t
	generate $t
	
done




echo "done"