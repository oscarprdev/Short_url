CREATE TABLE url_user (
    id UUID PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    url_id UUID NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE ON UPDATE CASCADE
);