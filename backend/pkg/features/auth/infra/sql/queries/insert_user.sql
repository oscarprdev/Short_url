INSERT INTO users (id, created_at, updated_at, email, picture, name, token, expires_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at, updated_at, email, picture, name, token, expires_at;