DROP INDEX IF EXISTS idx_pull_requests_status;
DROP INDEX IF EXISTS idx_pull_requests_author_id;

-- DROP в порядке зависимостией
DROP TABLE IF EXISTS pull_request_reviewers;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS team_members;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS teams;

DROP TYPE IF EXISTS pull_request_status;
