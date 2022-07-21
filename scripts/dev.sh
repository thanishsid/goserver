#!/bin/sh

echo "Starting Development Mode..."

docker run -p 5432:5432 -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=postgres \
    --name postgres_db -d --rm postgres:14-alpine3.16

docker run -p 6789:6789 --name redis_session -d --rm redis:7.0.3-alpine3.16

docker run -p 7700:7700 -e development=true --name ml_search -d --rm getmeili/meilisearch:latest

scripts/wait-for.sh localhost:5432 -- \
    scripts/wait-for.sh localhost:6379 -- \ 
scripts/wait-for.sh localhost:7700 -- air

docker stop postgres_db

docker stop redis_session

docker stop ml_search
