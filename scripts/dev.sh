#!/bin/sh

echo "Starting Development Mode..."

docker run -p 5432:5432 -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=postgres \
    --name postgres_db -d --rm postgres:14-alpine3.16

docker run -p 6379:6379 --name redis_session -d --rm redis:7.0.3-alpine3.16

scripts/wait-for.sh localhost:5432 -- \
    scripts/wait-for.sh localhost:6379 -- air

docker stop postgres_db

docker stop redis_session
