create table products
(
    id serial,
    name text not null,
    price numeric(10,2) not null default 0.00,
    constraint products_pkey primary key (id)
)
