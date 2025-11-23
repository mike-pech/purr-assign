package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mike-pech/purr-assign/cmd/api"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Server struct {
	repository Repository
}

func NewServer(r Repository) Server {
	return Server{
		repository: r,
	}
}

func (s Server) PostPullRequestCreate(ctx echo.Context) error {
	var pr api.PullRequest

	err := ctx.Bind(&pr)
	if err != nil {
		e := HTTPError{
			Code: http.StatusBadRequest,
			// TODO: Заменить сообщения ошибок
			Message: "Error in PostPullRequestCreate: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	newPR, err := s.repository.CreatePullRequesst(&pr)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error in CreatePullRequest: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}
	newPR, err = s.repository.AssignPullRequest(newPR)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error in AssignPullRequest: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusCreated, newPR)
}

func (s Server) PostPullRequestMerge(ctx echo.Context) error {
	id := ctx.Param("pull_request_id")
	pr, err := s.repository.SetPullRequestMerged(id)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusNotFound,
			Message: "Error in SetPullRequestMerged: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusOK, pr)
}

func (s Server) PostPullRequestReassign(ctx echo.Context) error {
	id := ctx.Param("pull_request_id")
	user := ctx.Param("old_user_id")

	pr, err := s.repository.ReassignPullRequest(id, user)
	if err != nil {
		switch err.Error() {
		case string(api.PRMERGED):
			e := HTTPError{
				Code:    http.StatusConflict,
				Message: "Error in ReassignPullRequest: " + err.Error(),
			}
			return ctx.JSON(e.Code, &e)
		case string(api.NOTASSIGNED):
			e := HTTPError{
				Code:    http.StatusConflict,
				Message: "Error in ReassignPullRequest: " + err.Error(),
			}
			return ctx.JSON(e.Code, &e)
		case string(api.NOCANDIDATE):
			e := HTTPError{
				Code:    http.StatusConflict,
				Message: "Error in ReassignPullRequest: " + err.Error(),
			}
			return ctx.JSON(e.Code, &e)
		}
		e := HTTPError{
			Code:    http.StatusNotFound,
			Message: "Error in ReassignPullRequest: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusOK, pr)
}

func (s Server) PostTeamAdd(ctx echo.Context) error {
	var team api.Team

	err := ctx.Bind(&team)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Error in PostTeamAdd: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	newTeam, err := s.repository.SetTeam(&team)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error in CreateTeam: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusCreated, newTeam)
}

func (s Server) GetTeamGet(ctx echo.Context, params api.GetTeamGetParams) error {
	query := params.TeamName
	teams, err := s.repository.FindTeam(query)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusNotFound,
			Message: "Error in FindTeam: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusOK, teams)
}

func (s Server) GetUsersGetReview(ctx echo.Context, params api.GetUsersGetReviewParams) error {
	id := params.UserId

	prs, err := s.repository.FindPullRequestsOfUser(&id)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error in FindPullRequestsOfUser: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusOK, prs)
}

func (s Server) PostUsersSetIsActive(ctx echo.Context) error {
	var setIsActiveBody api.PostUsersSetIsActiveJSONRequestBody

	err := ctx.Bind(&setIsActiveBody)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Error in PostUsersSetIsActive: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	user, err := s.repository.SetUserIsActive(&setIsActiveBody)
	if err != nil {
		e := HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Error in SetUserIsActive: " + err.Error(),
		}
		return ctx.JSON(e.Code, &e)
	}

	return ctx.JSON(http.StatusOK, user)
}
