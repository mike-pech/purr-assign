package models

import (
	"time"

	"github.com/mike-pech/purr-assign/cmd/api"
)

type PullRequest struct {
	// AssignedReviewers user_id назначенных ревьюверов (0..2)
	PullRequestId     string                `json:"pull_request_id" bun:"id,pk"`
	PullRequestName   string                `json:"pull_request_name" bun:"name,notnull"`
	AuthorId          *User                 `json:"author_id" bun:"author_id,rel:belongs-to"`
	AssignedReviewers []string              `json:"assigned_reviewers" bun:"rel:has-many,join:id=users_id"`
	CreatedAt         *time.Time            `json:"createdAt" bun:"created_at,nullzero,notnull,default:current_timestamp"`
	MergedAt          *time.Time            `json:"mergedAt" bun:"merged_at,nullzero,notnull,default:current_timestamp"`
	Status            api.PullRequestStatus `json:"status" bun:"status,notnull"`
}

type Team struct {
	TeamName string       `json:"team_name" bun:"name,pk"`
	Members  []TeamMember `json:"members" bun:"rel:has-many,join:id=user_id"`
}

type TeamMember struct {
	UserId   string `json:"user_id" bun:"user_id,pk"`
	Username string `json:"username" bun:"username,notnull,unique"`
	IsActive bool   `json:"is_active" bun:"is_active,notnull,default:true"`
}

type User struct {
	UserId   string `json:"user_id" bun:"id,pk"`
	Username string `json:"username" bun:"username,unique,notnull"`
	TeamName *Team  `json:"team_name" bun:"team,rel:belongs-to"`
	IsActive bool   `json:"is_active" bun:"is_active,notnull,default:true"`
}
