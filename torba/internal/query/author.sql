-- name: GetAuthorByID :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: GetAuthorByName :one
SELECT * FROM authors
WHERE nick_name = $1 LIMIT 1;


-- name: ListAuthorID :many
SELECT * FROM authors
ORDER BY id;

-- name: ListAuthorName :many
SELECT * FROM authors
ORDER BY nick_name;

-- name: SignAuthor :one
INSERT INTO authors (
    nick_name, email, password_hash
) VALUES (
             $1, $2, $3
         )
RETURNING id, nick_name;

-- name: SignFullAuthor :one
insert into authors(
    nick_name, email,
    password_hash,
    payments, projects,
    bio, link,
    profile_image_url, additional_info
)
values (
           $1, $2, $3,$4,$5,$6,$7,$8,$9
       )
returning *;

-- name: UpdateAuthor :one
UPDATE authors
set nick_name = $2,
    email = $3,
    password_hash = $4
WHERE id = $1
returning id,nick_name;

-- name: UpdateAuthorFull :one
UPDATE authors
set nick_name = $2,
    email = $3,
    password_hash = $4,
    payments = $5,
    projects = $6,
    bio = $7,
    link = $8,
    profile_image_url = $9,
    additional_info = $10
WHERE id = $1
returning id,nick_name;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;