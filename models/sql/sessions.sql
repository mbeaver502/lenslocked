create table sessions (
    id serial primary key,
    user_id int unique not null,
    token_hash text unique not null,

    foreign key (user_id) references users (id)
);

-- alter table sessions add constraint fk_sessions_user_id foreign key (user_id) references users (id);