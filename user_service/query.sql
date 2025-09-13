-- name: CreateUser :one
insert into users (email, password)
values (?, ?)
returning id, email;

-- name: GetUserByEmail :one
select id, email, password, role
from users
where email = ?;

-- name: CreateAdminUser :one
insert into users(email, password, role)
select ?, ?, 'admin'
where not exists (
  select 1 from users where role = 'admin'
)
returning id, email, role;
