-- name: GetTimeSchema :one
SELECT * FROM timeschemas
WHERE id = $1;

-- name: SearchTimeSchema :many
select * from timeschemas
where name like $1;

-- name: CreateTimeSchema :one
insert into timeschemas (name, items, updated_at)
values ($1, $2, now())
returning *;

-- name: UpdateTimeSchema :one
UPDATE timeschemas SET name = $2, items = $3, updated_at = now()
WHERE id = $1
returning *;

-- name: DeleteTimeSchema :one
delete from timeschemas where id = $1
returning *;
