SELECT urls.id, urls.original_url, urls.short_url, urls.title, urls.expires_at, urls.usage
FROM users
INNER JOIN url_user ON users.id = url_user.user_id
INNER JOIN urls ON url_user.url_id = urls.id
WHERE users.id = $1;