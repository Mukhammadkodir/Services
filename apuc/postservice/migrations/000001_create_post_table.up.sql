CREATE TABLE IF NOT EXISTS postos (
    id uuid primary key,
    userid TEXT,
    title varchar(90),
    images TEXT[],
    created_at TIMESTAMP,
    update_at TIMESTAMP,
    deleted_at TIMESTAMP
);