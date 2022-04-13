-- //Create comment
CREATE TABLE IF NOT EXISTS comment (
    id uuid primary key,
    user_id text,
    post_id text,
    text TEXT,
    created_at date,
    updated_at date,
    deleted_at date
);

                      