INSERT INTO url_user (id, user_id, url_id)
VALUES ($1, $2, $3) RETURNING *;