create table roles(
                      id serial primary key ,
                      name text not null unique ,
                      created_at timestamp default now(),
                      updated_at timestamp default now(),
                      description text
);

create table person_roles(
                           id serial primary key,
                           user_id integer not null,
                           role_id integer not null ,
                           assigned_at timestamp default now(),
                           foreign key (user_id) references person(id),
                           foreign key (role_id) references roles(id)
);

select *
from person;