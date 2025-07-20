-- +goose up
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- +goose down
ALTER TABLE feeds DROP COLUMN last_fetched_at;