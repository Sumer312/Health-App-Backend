-- name: CreateTotalCalorieIntake :one
INSERT INTO total_calorie_intake(id, user_id, calories, program, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
 
-- name: GetTotalCalories :many
SELECT * FROM total_calorie_intake WHERE user_id = $1;
 
-- name: GetMostRecentUserKcal :one
SELECT * FROM total_calorie_intake WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1;


-- name: DeleteRedundantRows :exec
DELETE FROM total_calorie_intake WHERE calories = 0;
