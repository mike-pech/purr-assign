package models

import "github.com/mike-pech/purr-assign/cmd/api/v1"

func (t *Team) ToAPI() *api.Team {
	var dtoMembers []api.TeamMember
	for _, m := range t.Members {
		dtoMembers = append(dtoMembers, api.TeamMember{
			UserId:   m.UserId,
			Username: m.Username,
			IsActive: m.IsActive,
		})
	}

	return &api.Team{TeamName: t.TeamName, Members: dtoMembers}
}

func TeamFromAPI(apiTeam *api.Team) *Team {
	var members []TeamMember
	for _, m := range apiTeam.Members {
		members = append(members, TeamMember{
			UserId:   m.UserId,
			Username: m.Username,
			IsActive: m.IsActive,
		})
	}

	return &Team{
		TeamName: apiTeam.TeamName,
		Members:  members,
	}
}

func (pr *PullRequest) ToAPI() *api.PullRequest {
	return &api.PullRequest{
		PullRequestId:     pr.PullRequestId,
		PullRequestName:   pr.PullRequestName,
		AuthorId:          pr.AuthorId,
		AssignedReviewers: pr.AssignedReviewers,
		CreatedAt:         pr.CreatedAt,
		MergedAt:          pr.MergedAt,
		Status:            pr.Status,
	}
}

func (pr *PullRequest) ToAPIShort() *api.PullRequestShort {
	return &api.PullRequestShort{
		PullRequestId:   pr.PullRequestId,
		PullRequestName: pr.PullRequestName,
		AuthorId:        pr.AuthorId,
		Status:          api.PullRequestShortStatus(pr.Status),
	}
}

func PullRequestFromAPI(apiPR *api.PullRequest, reviewers []string) *PullRequest {
	return &PullRequest{
		PullRequestId:     apiPR.PullRequestId,
		PullRequestName:   apiPR.PullRequestName,
		AuthorId:          apiPR.AuthorId,
		AssignedReviewers: reviewers,
		CreatedAt:         apiPR.CreatedAt,
		MergedAt:          apiPR.MergedAt,
		Status:            apiPR.Status,
	}
}

func (u *User) ToAPI() *api.User {
	return &api.User{
		UserId:   u.UserId,
		Username: u.Username,
		IsActive: u.IsActive,
	}
}
