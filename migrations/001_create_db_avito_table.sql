-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams (
    team_name TEXT PRIMARY KEY
);

CREATE TABLE users (
    user_id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    team_name TEXT REFERENCES teams(team_name),
    is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE pull_requests (
    pull_request_id TEXT PRIMARY KEY,
    pull_request_name TEXT NOT NULL,
    author_id TEXT REFERENCES users(user_id),
    status TEXT NOT NULL,
    merged_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE pull_request_reviewers (
    pr_id TEXT REFERENCES pull_requests(pull_request_id),
    reviewer_id TEXT REFERENCES users(user_id),
    PRIMARY KEY (pr_id, reviewer_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pull_request_reviewers;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;
-- +goose StatementEnd