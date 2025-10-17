-- +goose Up
-- goose StatementBegin

CREATE TABLE IF NOT EXISTS workouts_entries (
    id BIGSERIAL PRIMARY KEY,
    workouts_id BIGINT NOT NULL REFERENCES workouts(id) ON DELETE CASCADE,
    exercise_name VARCHAR(200) NOT NULL,
    sets INT NOT NULL,
    reps INT,
    duration_seconds INT,
    weight DECIMAL(5,2),
    notes TEXT,
    order_index INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT vailid_workout_entry CHECK (
        ( reps IS NOT NULL OR duration_seconds IS NOT NULL) AND
        ( reps IS NULL OR duration_seconds is NULL)
    )
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workouts_entries;
