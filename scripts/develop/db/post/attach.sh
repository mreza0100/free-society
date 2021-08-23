#!/bin/bash
clear




docker container exec -ti docker_post_service_postgres_1 /bin/bash -c 'psql -U postgres -p 5434'



