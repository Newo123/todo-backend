CREATE SCHEMA todo;

CREATE TABLE todo.users (
    id UUID PRIMARY KEY,
    version BIGINT NOT NULL,
    full_name VARCHAR(100),
    password_hash TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE todo.tasks (
    id UUID PRIMARY KEY,
    version BIGINT NOT NULL,
    title VARCHAR(100) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 100),
    description VARCHAR(1000) CHECK(char_length(description) BETWEEN 1 AND 1000),
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    completed_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL,

    CHECK(
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    ),

    user_id UUID NOT NULL REFERENCES todo.users(id) ON DELETE CASCADE
);

CREATE INDEX idx_tasks_author_user_id ON todo.tasks (user_id);
CREATE INDEX idx_tasks_user_completed ON todo.tasks(user_id, completed);