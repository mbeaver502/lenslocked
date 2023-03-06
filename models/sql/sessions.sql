CREATE TABLE sessions (
    id serial PRIMARY KEY,
    user_id int UNIQUE NOT NULL,
    token_hash text UNIQUE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- alter table sessions add constraint fk_sessions_user_id foreign key (user_id) references users (id);