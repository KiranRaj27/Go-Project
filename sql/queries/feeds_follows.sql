-- name: CreateFeedFollows :one
INSERT INTO feed_follows(id,created_at,updated_at,user_id,feed_id)
VALUES($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id=$1;



