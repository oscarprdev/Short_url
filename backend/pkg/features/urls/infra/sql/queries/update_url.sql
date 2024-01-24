UPDATE urls
SET usage = $1, updated_at = $2
WHERE id = $3 
RETURNING *;