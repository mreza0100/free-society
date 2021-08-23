#!/bin/bash
clear




# psql --host localhost --user postgres --port 5437
docker container exec -ti docker_notification_service_postgres_1 /bin/bash -c 'psql -U postgres -p 5437'




