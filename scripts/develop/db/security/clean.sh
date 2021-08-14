#!/bin/bash




psql --host localhost --user postgres --port 5436 --command="DROP TABLE sessions"
psql --host localhost --user postgres --port 5436 --command="DROP TABLE passwords"



echo "flushall" | redis-cli -h localhost -p 6380

