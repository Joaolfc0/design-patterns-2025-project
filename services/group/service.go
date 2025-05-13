package group

import (
	"math/rand"
	"service-secret-santa/customError"
	"service-secret-santa/functions"
	"service-secret-santa/models"
	"service-secret-santa/repositories/group"
	"time"
)

type Service interface {
	CreateGroup(group *models.Group) (*models.Group, *customError.CustomError)
	GetGroupByID(id string) (*models.Group, *customError.CustomError)
	UpdateGroup(id string, group *models.Group) (*models.Group, *customError.CustomError)
	DeleteGroup(id string) *customError.CustomError
	AddParticipant(id string, participant *models.Participant) (*models.Group, *customError.CustomError)
	MatchParticipants(id string) (*models.Group, *customError.CustomError)
	GetMyMatch(id string, username string) (string, *customError.CustomError)
	GetAllGroups() ([]*models.Group, *customError.CustomError)
}

type resource struct {
	repo group.Repository
}

func (r *resource) CreateGroup(group *models.Group) (*models.Group, *customError.CustomError) {
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	return r.repo.CreateGroup(group)
}

func (r *resource) GetGroupByID(id string) (*models.Group, *customError.CustomError) {
	return r.repo.GetGroupByID(id)
}

func (r *resource) UpdateGroup(id string, group *models.Group) (*models.Group, *customError.CustomError) {
	group.UpdatedAt = time.Now()
	return r.repo.UpdateGroup(id, group)
}

func (r *resource) DeleteGroup(id string) *customError.CustomError {
	return r.repo.DeleteGroup(id)
}

func (r *resource) AddParticipant(id string, participant *models.Participant) (*models.Group, *customError.CustomError) {
	return r.repo.AddParticipant(id, participant)
}

func (r *resource) MatchParticipants(id string) (*models.Group, *customError.CustomError) {
	group, err := r.repo.GetGroupByID(id)
	if err != nil {
		return nil, err
	}

	if len(group.Participants) < 2 {
		return nil, customError.CustomErrorFactory(customError.WithBadRequest("Not enough participants", "At least two participants are required for matching"))
	}

	// Maps names of participants
	remainingParticipants := make(map[string]struct{})
	for _, participant := range group.Participants {
		remainingParticipants[participant.Name] = struct{}{}
	}

	var matches []models.Match
	matchIndexes := functions.RandomDerangement(len(group.Participants), rand.New(rand.NewSource(time.Now().UnixNano())))

	for i, participant := range group.Participants {
		matches = append(matches, models.Match{
			First:  participant.Name,
			Second: group.Participants[matchIndexes[i]].Name,
		})
	}

	// Updates matches in repo
	updateErr := r.repo.UpdateMatches(id, matches)
	if updateErr != nil {
		return nil, updateErr
	}

	// Updates matches in group
	group.Matches = matches
	group.NotifyAll("matchParticipants")

	return group, nil
}

func (r *resource) GetMyMatch(id string, username string) (string, *customError.CustomError) {
	return r.repo.GetMyMatch(id, username)
}

func (r *resource) GetAllGroups() ([]*models.Group, *customError.CustomError) {
	return r.repo.GetAllGroups()
}

func NewGroupService(repo group.Repository) Service {
	return &resource{repo: repo}
}
