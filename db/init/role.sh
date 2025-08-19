#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 \
    --username "$DB_USER" \
    --dbname "$DB_NAME" <<-EOSQL
    CREATE ROLE $APP_DB_USER WITH LOGIN PASSWORD '$APP_DB_PASSWORD' NOSUPERUSER;
    CREATE DATABASE $APP_DB_NAME OWNER $APP_DB_USER;
    ALTER ROLE $APP_DB_USER SET search_path = app, public;
EOSQL