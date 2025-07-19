-- +goose up
CREATE TABLE feeds_followers (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    user_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id  UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);

-- +goose down
DROP TABLE feeds_followers;
