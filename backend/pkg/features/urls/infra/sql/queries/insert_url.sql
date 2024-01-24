INSERT INTO urls (id, created_at, updated_at, original_url, short_url, title, expires_at, usage)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;