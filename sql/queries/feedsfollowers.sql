-- name: CreateFeedsFollowers :many
INSERT INTO feeds_followers(id, created_at, updated_at,user_id,feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: CreateFeedFollow :one
WITH inserted AS (
    INSERT INTO feeds_followers (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    inserted.id,
    inserted.created_at,
    inserted.updated_at,
    inserted.user_id,
    inserted.feed_id,
    users.name AS user_name,
    feeds.name AS feed_name,
    feeds.url AS feed_url
FROM inserted
JOIN users ON inserted.user_id = users.id
JOIN feeds ON inserted.feed_id = feeds.id;

 
-- name: GetFeedFollowsForUser :many
SELECT 
    ff.id,
    ff.created_at,
    ff.updated_at,
    ff.user_id,
    ff.feed_id,
    u.name AS user_name,
    f.name AS feed_name,
    f.url AS feed_url
FROM feeds_followers ff
JOIN users u ON ff.user_id = u.id
JOIN feeds f ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: UnfollowFeed :exec
DELETE FROM feeds_followers ff
WHERE ff.user_id = $1 
AND ff.feed_id = (
    SELECT id FROM feeds WHERE url = $2
);