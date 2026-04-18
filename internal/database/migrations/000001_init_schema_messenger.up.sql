create table messenger_users (
    user_id uuid primary key,
    user_name varchar(15) not null,
    user_login varchar(10) not null,
    user_password varchar(10) not null,
    user_email varchar(50) not null,
    created_at timestamp default now(),
);