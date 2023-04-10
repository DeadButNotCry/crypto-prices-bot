CREATE DATABASE crypto_prices;
CREATE USER bnc WITH PASSWORD 'changeme';
GRANT ALL PRIVILEGES ON DATABASE "crypto_prices" TO bnc;

CREATE TABLE student
(
    id bigint NOT NULL,
    name text COLLATE pg_catalog."default",
    CONSTRAINT student_pkey PRIMARY KEY (id)
);

INSERT INTO student(id, name) VALUES
                                  (1, 'A'),
                                  (2, 'B'),
                                  (3, 'C');