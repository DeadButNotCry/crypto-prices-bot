FROM postgres:latest

ADD ./migration_01.sql /migrations/
ADD ./init-db.sh /docker-entrypoint-initdb.d/
