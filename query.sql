-- name: getUser :one
SELECT user_id, email, password FROM user
WHERE id = ? LIMIT 1;

-- name: getUserAuth :one
SELECT u.user_id, u.email, u.password, a.password_hash, a.session_token, a.csrf_token
FROM users u INNER JOIN auth a 
ON u.user_id = a.user_id
WHERE u.user_id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  email, password
) VALUES (
  ?, ?
);

-- name: UpdateUser :execresult
UPDATE users 
SET email = ?, password = ?
WHERE user_id = ?;

-- name: UpdateUserAuth :execresult
UPDATE auth 
SET session_token= ?, password_hash = ?, csrf_token = ?
WHERE auth_id = ? OR user_id = ?;

-- name: DeleteUserAuth :exec
DELETE FROM auth
WHERE auth_id = ? OR user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;
