create table if not exists users (
    id serial primary key,
    email text not null unique,
    password text not null,
    created_at timestamp not null default current_timestamp
);

create unique index users_email_unique on users (email);