-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUserID :many
SELECT * FROM users
ORDER BY id;

-- name: ListUserName :many
SELECT * FROM users
ORDER BY username;

-- name: SignUser :one
INSERT INTO users (
    username, email, password_hash
) VALUES (
             $1, $2, $3
         )
RETURNING id,username;

-- name: SignFullUser :one
insert into users(
    username, email, password_hash,
    donation_sum, supported_projects,
    profile_image_url
)
values (
        $1, $2, $3,$4,$5,$6
)
returning *;

-- name: UpdateUser :one
UPDATE users
set username = $2,
    email = $3,
    password_hash = $4
WHERE id = $1
returning id,username;

-- name: UpdateFull :one
UPDATE users
set username = $2,
    email = $3,
    password_hash = $4,
    donation_sum = $5,
    supported_projects = $6,
    profile_image_url = $7
WHERE id = $1
returning id,username;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;