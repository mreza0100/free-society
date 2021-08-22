#!/bin/bash
clear

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $DIR
cd ../..
ROOT=$(pwd)



cd $ROOT/scripts/develop/db/
dbs=$(ls -d ./*/)
echo $dbs

for t in ${dbs[@]};
do
      echo 'Running' $t
      bash $ROOT/scripts/develop/db/$t/start.sh --detach
      echo $t 'Done'
      echo '--------------------------------------------------------------------'
done

bash $ROOT/scripts/develop/common/start.sh --detach

exit 0