package router

import (
	"crusadetrackerapi/internal/armies"
	"crusadetrackerapi/internal/factions"
	"crusadetrackerapi/internal/rosters"
	"crusadetrackerapi/internal/users"
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func StartServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONNSTRING"))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	armyService := armies.NewService(armies.NewRepository(db))
	router := gin.Default()

	armyService.RegisterRoutes(router)
	factions.RegisterRoutes(router)
	users.RegisterRoutes(router)
	rosters.RegisterRoutes(router)

	router.Run()
}
