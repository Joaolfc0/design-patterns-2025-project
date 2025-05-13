package group

import (
	groupHandler "service-secret-santa/handlers/group"

	"github.com/gin-gonic/gin"
)

// Routes sets up the routes for the group resource
func Routes(defaultGroup *gin.RouterGroup, handler groupHandler.Handler) {
	groupsGroup := defaultGroup.Group("/group")
	{
		// default group = http://localhost:8080/secret-santa/

		// Route to create a group
		groupsGroup.POST("", handler.CreateGroup)

		// Route to get a group by ID
		groupsGroup.GET("/:id", handler.GetGroup)

		// Route to update a group by ID
		groupsGroup.PUT("/:id", handler.UpdateGroup)

		// Route to delete a group by ID
		groupsGroup.DELETE("/:id", handler.DeleteGroup)

		// Route to add a participant to the group
		groupsGroup.POST("/:id/add-participant", handler.AddParticipant)

		// Route to generate the matches for the group's participants
		groupsGroup.POST("/:id/match-participants", handler.MatchParticipants)

		// Route to get the match of a participant
		groupsGroup.GET("/:id/my-match", handler.GetMyMatch)

		// Route to get all available groups
		groupsGroup.GET("", handler.GetAllGroups)
	}
}
