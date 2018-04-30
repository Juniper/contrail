
create table customers (
    id varchar(255),
    primary key(id)
);

create table books (
    id varchar(255),
    customer_id varchar(255),
    price integer,
    primary key(id),
    FOREIGN KEY (customer_id) REFERENCES customers(id)
);

insert into customers (id) values ('alice');
insert into books (id, customer_id, price) values ('apple', 'alice', 100);
insert into books (id, customer_id, price) values ('banana', 'alice', 200);
insert into books (id, customer_id, price) values ('orange', 'alice', 300);