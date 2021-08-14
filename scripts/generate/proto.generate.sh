#!/bin/bash 
clear




DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../
ROOT=$(pwd)

generated=$ROOT/proto/generated
raw_protos=$ROOT/proto/raw_protos


function generate {
	protoc --go_out=$generated/$1 --go_opt=paths=source_relative \
		--go-grpc_out=$generated/$1 --go-grpc_opt=paths=source_relative \
		./*.proto
}



protos=$(ls $ROOT/proto/raw_protos)


for t in ${protos[@]};
do
	echo $t
	cd $ROOT
	mkdir -p $generated/$t
	cd $raw_protos/$t
	generate $t
	echo "---"
done




echo "done"