-- name: GetAllProducts :many
select * from products
where isDeleted = false;
-- name: GetFilterdProducts :many
select * from products
where price >= :min 
and price <= :max
and isDeleted = false;

-- name: GetProductByID :one
select * from products
where id = ?
and isDeleted = false;

-- name: GetAllProductsForAdmin :many
select * from products;

-- name: CreateProduct :one
insert into products
(name, price, stock)
values (?, ?, ?)
returning *;

-- name: UpdateProductByID :one
update products
set name = ?, price = ?, stock = ?
where id = ?
returning *;

-- name: DeleteProductByID :one
update products
set isDeleted = true
where id = ?
returning *;
