#!/bin/bash
clear

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR/../../
ROOT=$(pwd)



DBS=$(ls $ROOT/scripts/develop/db)
for db in ${DBS[@]};
do
      echo "Cleaning $db"
      bash $ROOT/scripts/develop/db/$db/clean.sh
      echo "--------"
done


mv $ROOT/public/avatars/default_female.jpg $ROOT/public/avatars/default_male.jpeg $ROOT/public
rm -rf $ROOT/public/avatars/* $ROOT/public/images/*
mv $ROOT/public/default_female.jpg $ROOT/public/default_male.jpeg $ROOT/public/avatars
