#!/bin/bash

set -e

podman network create common-inv-db-network || true

podman run \
  --name=common-inv-db \
  --net common-inv-db-network \
  --restart=always \
  -e POSTGRES_PASSWORD=test \
  -e POSTGRES_USER=test \
  -e POSTGRES_DB=commondb \
  -p 5432:5432 \
  -d postgres:latest \
  postgres -N 200
