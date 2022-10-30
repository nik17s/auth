CREATE TABLE users
(
    id       BIGSERIAL PRIMARY KEY NOT NULL,
    login    VARCHAR(255) UNIQUE   NOT NULL,
    password VARCHAR(255)          NOT NULL,
    email    VARCHAR(255)          NOT NULL,
    phone    VARCHAR(255)          NOT NULL
);
