UPDATE users
SET token = $1, expires_at = $2
WHERE id = $3;