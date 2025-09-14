create table products (
id integer primary key autoincrement,
  name text not null,
  price numeric(10,2) not null check (price >=0),
  stock integer not null check (stock >=0),
  isDeleted boolean not null default false
);
