-- name: getUser :one
SELECT user_id, email, password FROM users
WHERE user_id = ? LIMIT 1;


-- name: CheckUserExists :one
SELECT EXISTS(
  SELECT 1 FROM users WHERE email = ?
);

-- name: getUserAuth :one
SELECT u.user_id, u.email, u.password, a.password_hash, a.session_token, a.csrf_token
FROM users u INNER JOIN auth a 
ON u.user_id = a.user_id
WHERE u.email = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (
  email, password
) VALUES (
  ?, ?
);

-- name: CreateAuth :execresult
INSERT INTO auth (
 user_id, password_hash 
) VALUES ( ?, ? );

-- name: GetAuth :one
SELECT a.auth_id, a.password_hash, a.session_token, a.csrf_token
FROM auth a INNER JOIN users u
ON a.user_id = u.user_id
WHERE u.email = ?;

-- name: RetrieveAuth :one
SELECT a.auth_id, a.password_hash, a.session_token, a.csrf_token
FROM auth a
WHERE a.csrf_token = ? AND a.session_token = ?;

-- name: UpdateUser :execresult
UPDATE users 
SET email = ?, password = ?
WHERE user_id = ?;

-- name: UpdateUserAuth :execresult
UPDATE auth 
SET session_token= ?, password_hash = ?, csrf_token = ?
WHERE auth_id = ?;

-- name: DeleteUserAuth :exec
DELETE FROM auth
WHERE auth_id = ? OR user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;
