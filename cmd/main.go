package main

import (
	"fmt"
	"log"
	"social-network/internal/config"
	categoryHandler "social-network/internal/handler/categories"
	userHandler "social-network/internal/handler/users"
	categoryRepo "social-network/internal/repository/categories"
	userRepo "social-network/internal/repository/users"
	categorySvc "social-network/internal/service/categories"
	userSvc "social-network/internal/service/users"
	"social-network/pkg/internalsql"

	"github.com/gin-gonic/gin"

	_ "time/tzdata"
)

func main() {
	r := gin.Default()

	// Load configuration
	cfg, err := config.SetupConfig()
	if err != nil {
		log.Println("failed load configuration: ", err)
	}

	// Load database connection
	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		log.Println("failed load database connection: ", err)
	}

	// Categories API
	categoryRepo := categoryRepo.NewRepository(db)
	categoryService := categorySvc.NewService(categoryRepo)
	categoryHandler := categoryHandler.NewHandler(r, categoryService)
	categoryHandler.RouteList()

	// Users API
	userRepo := userRepo.NewRepository(db)
	userService := userSvc.NewService(userRepo, cfg)
	userHandler := userHandler.NewHandler(r, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("127.0.0.1:%s", cfg.PORT)
	r.Run(server)
}
