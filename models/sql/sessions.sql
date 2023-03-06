create table sessions (
    id serial primary key,
    user_id int unique not null,
    token_hash text unique not null
);