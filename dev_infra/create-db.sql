create database kodingkorp;
\c kodingkorp;
create table users (id serial not null, email varchar(255) not null, name varchar(255) not null, created_at timestamp DEFAULT CURRENT_TIMESTAMP not null);
insert into users (email, name) values ('shivam@kodingkorp.com', 'Shivam Mathur')