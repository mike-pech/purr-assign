package repository

import (
	"context"
	"database/sql"

	"github.com/mike-pech/purr-assign/cmd/api/v1"
	"github.com/mike-pech/purr-assign/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type BunRepository struct {
	db *bun.DB
}

func NewBunRepository(conn *sql.DB) BunRepository {
	return BunRepository{
		db: bun.NewDB(conn, pgdialect.New()),
	}
}

func (r BunRepository) CreatePullRequest(ctx context.Context, apiPR *api.PullRequest) (*api.PullRequest, error) {
	var reviewers []string
	err := r.db.NewSelect().
		Model(&reviewers).
		Table("users").
		OrderExpr("RANDOM()").
		Where("is_active").
		Limit(2).
		Scan(ctx)

	pr := models.PullRequestFromAPI(apiPR, reviewers)

	_, err = r.db.NewInsert().
		Model(&pr).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return pr.ToAPI(), err
}

func (r BunRepository) SetPullRequestMerged(ctx context.Context, id string) (*api.PullRequest, error) {
	pr := models.PullRequest{
		PullRequestId: id,
		Status:        "MERGED",
	}

	_, err := r.db.NewUpdate().
		Model(&pr).
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return pr.ToAPI(), err
}

func (r BunRepository) ReassignPullRequest(ctx context.Context, id, user string) (*api.PullRequest, error) {
	var reviewers []string
	err := r.db.NewSelect().
		Model(&reviewers).
		Table("users").
		OrderExpr("RANDOM()").
		Where("is_active").
		Where("NOT id = ?", user).
		Limit(2).
		Scan(ctx)

	pr := models.PullRequest{
		PullRequestId:     id,
		AssignedReviewers: reviewers,
	}

	_, err = r.db.NewUpdate().
		Model(&pr).
		WherePK().
		Where("NOT status = MERGED").
		OmitZero().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return pr.ToAPI(), err
}

func (r BunRepository) SetTeam(ctx context.Context, team *api.Team) (*api.Team, error) {
	t := models.TeamFromAPI(team)

	_, err := r.db.NewInsert().
		Model(&t).
		On("CONFLICT (team_name) DO UPDATE").
		Set("name = ?", team.TeamName).
		Set("members = ?", team.Members).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return t.ToAPI(), err
}

func (r BunRepository) FindTeam(ctx context.Context, query string) (*[]api.Team, error) {
	var (
		apiTeams []api.Team
		teams    []models.Team
	)

	err := r.db.NewSelect().
		Model(&teams).
		Table("teams").
		Where("name = ?", query).
		WherePK().
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	for _, t := range teams {
		apiTeams = append(apiTeams, *t.ToAPI())
	}

	return &apiTeams, err
}

func (r BunRepository) FindPullRequestsOfUser(ctx context.Context, query *api.UserIdQuery) (*[]api.PullRequestShort, error) {
	var prs []models.PullRequest

	err := r.db.NewSelect().
		Model(&prs).
		Table("pull_requests").
		Where("user_id = ?", query).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	var apiPRs []api.PullRequestShort
	for _, p := range prs {
		apiPRs = append(apiPRs, *p.ToAPIShort())
	}

	return &apiPRs, err
}

func (r BunRepository) SetUserIsActive(ctx context.Context, userIsActiveRequest *api.PostUsersSetIsActiveJSONRequestBody) (*api.User, error) {
	u := models.User{
		UserId:   userIsActiveRequest.UserId,
		IsActive: userIsActiveRequest.IsActive,
	}

	_, err := r.db.NewUpdate().
		Model(&u).
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return u.ToAPI(), err
}
