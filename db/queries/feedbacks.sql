-- name: CreateFeedback :one
INSERT INTO feedbacks (email, message, user_id)
VALUES ($1, $2, $3)
RETURNING id;

-- name: SelectAllFeedbacks :many
SELECT * FROM feedbacks;

-- name: UpdateFeedback :one
UPDATE feedbacks
SET
    email = $2,
    message = $3
WHERE id = $1
RETURNING id;

-- name: DeleteFeedback :one
DELETE FROM feedbacks WHERE id = $1
RETURNING id;
