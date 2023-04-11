CREATE TABLE IF NOT EXISTS app_user
(
    id                BIGINT PRIMARY KEY,
    is_admin          bool         NOT NULL,
    is_banned         bool         NOT NULL,
    state             VARCHAR(100) NOT NULL,
    last_message_date TIMESTAMP    NOT NULL,
    register_date     DATE         NOT NULL
);