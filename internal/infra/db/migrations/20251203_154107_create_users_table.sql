-- Migration: create_users_table
-- Created at: 2025-12-03T15:41:07+01:00

CREATE TABLE IF NOT EXISTS users
(
    id       int PRIMARY KEY,
    name     VARCHAR(255)        NOT NULL,
    email    VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255)        NOT NULL
);
