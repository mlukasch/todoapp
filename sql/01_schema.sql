CREATE TABLE IF NOT EXISTS TODO_USER(
    id serial Primary Key,
    firstName varchar(255) not null,
    lastName varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    UNIQUE(email)
);