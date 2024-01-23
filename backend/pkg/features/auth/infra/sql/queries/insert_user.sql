INSERT INTO users (id, created_at, updated_at, email, picture, name) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, email, picture, name;