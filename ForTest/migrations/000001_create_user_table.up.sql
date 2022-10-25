CREATE TABLE IF NOT EXISTS testuser (
    id uuid primary key,
    name VARCHAR(64),
    username VARCHAR(32),
    city VARCHAR(64),
    created_at date,
    updated_at date,
    deleted_at date
);
