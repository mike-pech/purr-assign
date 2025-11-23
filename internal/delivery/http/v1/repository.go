package v1

import (
	"context"

	"github.com/mike-pech/purr-assign/cmd/api/v1"
)

type Repository interface {
	// Создать PR и назначить 1–2 ревьювера на PR
	CreatePullRequest(ctx context.Context, pr *api.PullRequest) (*api.PullRequest, error)
	// Пометить PR как MERGED (идемпотентная операция)
	SetPullRequestMerged(ctx context.Context, id string) (*api.PullRequest, error)
	// Переназначить конкретного ревьювера на другого из его команды
	ReassignPullRequest(ctx context.Context, id, user string) (*api.PullRequest, error)
	// Создать или обновить команду
	SetTeam(ctx context.Context, team *api.Team) (*api.Team, error)
	// Найти команды по запросу
	FindTeam(ctx context.Context, query string) (*[]api.Team, error)
	// Получить PR'ы, где пользователь назначен ревьювером
	FindPullRequestsOfUser(ctx context.Context, query *api.UserIdQuery) (*[]api.PullRequestShort, error)
	// Установить флаг активности пользователя
	SetUserIsActive(ctx context.Context, userIsActiveRequest *api.PostUsersSetIsActiveJSONRequestBody) (*api.User, error)
}
