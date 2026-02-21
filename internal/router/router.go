package router

import (
	"crusadetrackerapi/internal/armies"
	"crusadetrackerapi/internal/factions"
	"crusadetrackerapi/internal/rosters"
	"crusadetrackerapi/internal/users"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	armyService := armies.Service{}
	router := gin.Default()

	armyService.RegisterRoutes(router)
	factions.RegisterRoutes(router)
	users.RegisterRoutes(router)
	rosters.RegisterRoutes(router)

	router.Run()
}
