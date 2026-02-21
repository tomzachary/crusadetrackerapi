package router

import (
	"crusadetrackerapi/internal/armies"
	"crusadetrackerapi/internal/factions"
	"crusadetrackerapi/internal/rosters"
	"crusadetrackerapi/internal/users"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	armies.RegisterRoutes(router)
	factions.RegisterRoutes(router)
	users.RegisterRoutes(router)
	rosters.RegisterRoutes(router)

	router.Run()
}
