#!/bin/sh

 psql -v --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -a -f /migrations/migration_01.sql