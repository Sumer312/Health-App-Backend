-- +goose Up 

CREATE TABLE user_input (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  program VARCHAR(16) NOT NULL,
  sex VARCHAR(6) NOT NULl,
  height INT NOT NULL,
  weight INT NOT NULL,
  desired_weight INT,
  time_frame INT,
  bmi FLOAT NOT NULL,
  curr_kcal FLOAT NOT NULl,
  deficit FLOAT
);

-- +goose Down
DROP TABLE user_input;

