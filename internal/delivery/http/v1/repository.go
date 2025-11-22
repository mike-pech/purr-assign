package v1

import "github.com/mike-pech/purr-assign/cmd/api"

type Repository interface {
	// Создать PR
	CreatePullRequesst(*api.PullRequest) (*api.PullRequest, error)
	// Назначить ревьювера на PR
	AssignPullRequest(*api.PullRequest) (*api.PullRequest, error)
	// Пометить PR как MERGED (идемпотентная операция)
	SetPullRequestMerged(pr string) (*api.PullRequest, error)
	// Переназначить конкретного ревьювера на другого из его команды
	ReassignPullRequest(pr, user string) (*api.PullRequest, error)
	// Создать или обновить команду
	SetTeam(*api.Team) (*api.Team, error)
	// Найти команды по запросу
	FindTeam(query string) (*[]api.Team, error)
	// Получить PR'ы, где пользователь назначен ревьювером
	FindPullRequestsOfUser(*api.UserIdQuery) (*[]api.PullRequestShort, error)
	// Установить флаг активности пользователя
	SetUserIsActive(*api.PostUsersSetIsActiveJSONRequestBody) (*api.User, error)
}
