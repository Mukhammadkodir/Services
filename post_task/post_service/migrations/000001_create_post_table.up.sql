-- //Create post
CREATE TABLE IF NOT EXISTS post7 (
    id VARCHAR(10),
    user_id VARCHAR(10),
    title TEXT,
    body TEXT,
    created_at date,
    updated_at date,
    deleted_at date
);


alter TABLE migrat add column if not EXISTS newcolumn TEXT;
alter TABLE if EXISTS migrat drop column if EXISTS columnname;
