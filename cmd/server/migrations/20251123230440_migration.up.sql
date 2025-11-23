CREATE TYPE pull_request_status AS ENUM ('OPEN', 'CLOSED', 'MERGED', 'DRAFT');

CREATE TABLE teams (
	name VARCHAR(50) PRIMARY KEY,
	members VARCHAR(50) ARRAY
);

CREATE TABLE users (
	user_id VARCHAR(50) PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	team_name VARCHAR(50),
	is_active BOOLEAN NOT NULL DEFAULT true,

	FOREIGN KEY (team_name) REFERENCES teams(name) ON DELETE SET NULL
);

CREATE TABLE team_members (
	user_id VARCHAR(50) PRIMARY KEY,
	username VARCHAR(255) UNIQUE NOT NULL,
	is_active BOOLEAN NOT NULL DEFAULT true,

	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
);

CREATE TABLE pull_requests (
	id VARCHAR(50) PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	author_id VARCHAR(50) NOT NULL,
	assigned_reviewers VARCHAR(50) ARRAY,
	created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	merged_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	status pull_request_status NOT NULL,

	FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_pull_requests_author_id ON pull_requests(author_id);
CREATE INDEX idx_pull_requests_status ON pull_requests(status);
