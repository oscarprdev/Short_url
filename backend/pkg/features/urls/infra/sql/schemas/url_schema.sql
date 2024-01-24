DROP TABLE urls IF EXIST;

CREATE TABLE urls (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    original_url VARCHAR(255) NOT NULL UNIQUE,
    short_url TEXT NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 week'
);

CREATE TRIGGER update_expires_at
AFTER UPDATE
ON url_statistics FOR EACH ROW
BEGIN
    IF NEW.usage_count > OLD.usage_count THEN
        UPDATE urls
        SET expires_at = DATE_ADD(NEW.access_timestamp, INTERVAL 1 DAY)
        WHERE id = NEW.url_id;
    END IF;
END;