#!/bin/bash




psql --host localhost --user postgres --port 5435 --command="DROP TABLE followers"
psql --host localhost --user postgres --port 5435 --command="DROP TABLE likes"


