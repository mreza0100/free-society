#!/bin/bash
clear

dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cd $dir
cd ../../../..
root=$(pwd)



psql --host localhost --user postgres --port 5433 -f "$root/services/user/db/init.sql"


