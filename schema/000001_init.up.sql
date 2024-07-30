CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE tasks
(
    id        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title     VARCHAR(255) NOT NULL,
    active_at date,
    status    BOOLEAN DEFAULT FALSE
);
