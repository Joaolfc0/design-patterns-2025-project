package models

import (
	"fmt"
	"time"

	"github.com/invopop/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Group struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty" swaggerignore:"true"`
	Name         string             `json:"name" bson:"name" example:"Equipe pe no chao"`
	Participants []Participant      `json:"participants" bson:"participants" `
	Observers    []Observer         `json:"-" bson:"-" swaggerignore:"true"`
	Matches      []Match            `json:"matches" bson:"matches"  swaggerignore:"true"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt, omitempty" swaggerignore:"true"`
	UpdatedAt    time.Time          `json:"updateAt" bson:"updateAt, omitempty" swaggerignore:"true"`
}

var mockGroupID = func() primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex("6787c4a755ea623ab45e77d4")
	return id
}()

func CreateMockGroup() *Group {
	return &Group{
		Id:           mockGroupID,
		Name:         "Test Group",
		Participants: []Participant{{Name: "Mari", Email: "mari@gmail.com"}},
		Matches:      []Match{{First: "joao", Second: "mari"}},
		CreatedAt:    time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2023, 12, 11, 0, 0, 0, 0, time.UTC),
	}
}

type Participant struct {
	Name  string `json:"name" bson:"name" example:"Mari"`
	Email string `json:"email" bson:"email" example:"Mari@gmail.com"`
}

// Update satisfies the Observer interface
func (p Participant) Update(eventType string, group *Group) {
	if eventType == "matchParticipants" {
		for _, match := range group.Matches {
			if match.First == p.Name {
				fmt.Printf("Sending email to %s: You are gifting %s!\n", p.Email, match.Second)
				return
			}
		}
	}
}

type Match struct {
	First  string `json:"first" bson:"first"  example:"joao"`
	Second string `json:"second" bson:"second" example:"mari"`
}

func (l Group) Validate() error {
	err := validation.ValidateStruct(&l,
		validation.Field(&l.Name, validation.Required),
	)

	if err != nil {
		return err
	}

	return nil
}

// NotifyAll notifies all participants about the event
func (g *Group) NotifyAll(eventType string) {
	for _, observer := range g.Observers {
		observer.Update(eventType, g)
	}
}

func (g *Group) RebuildObservers() {
	g.Observers = []Observer{}
	for _, participant := range g.Participants {
		g.Observers = append(g.Observers, participant)
	}
}

func (l Participant) Validate() error {
	err := validation.ValidateStruct(&l,
		validation.Field(&l.Name, validation.Required),
		validation.Field(&l.Email, validation.Required),
	)

	if err != nil {
		return err
	}

	return nil
}
