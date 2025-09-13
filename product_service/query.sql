-- name: GetProducts :many
-- go-type: min=float64
-- go-type: max=float64
select * from products
where isDeleted = false;
-- name: GetFilteredProducts :many
select * from products
where price >= :min 
and price <= :max
and isDeleted = false;

-- name: GetProductByID :one
select * from products
where id = ?
and isDeleted = false;

-- name: GetProductsForAdmin :many
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
