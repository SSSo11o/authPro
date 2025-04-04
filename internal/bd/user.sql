create table persons (
    id serial primary key ,
    name text not null ,
    email text unique not null ,
    password text not null ,
    is_active boolean default true,
    created_at timestamp default now(),
    updated_at timestamp default now()
);





