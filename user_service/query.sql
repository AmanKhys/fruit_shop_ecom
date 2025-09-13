-- name: CreateUser :one
insert into users (email, password)
values (?, ?)
returning id, email;

-- name: GetUserByEmail :one
select id, email, password
from users
where email = ?;
