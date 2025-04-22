-- name: getUser :one
SELECT user_id, email, password FROM user
WHERE id = ? LIMIT 1;

-- name: getUserAuth :one
SELECT u.user_id, u.email, u.password, a.password_hash, a.session_token, a.csrf_token
FROM users u INNER JOIN auth a 
ON u.user_id = a.user_id
WHERE u.user_id = ? LIMIT 1;

-- update
