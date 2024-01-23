INSERT INTO urls (id, original_url, short_url)
VALUES ($1, $2, $3) RETURNING *;