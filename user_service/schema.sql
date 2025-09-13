create table users (
  id integer primary key autoincrement,
  email text not null unique,
  password text not null
);
