package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/panjiasmoroart/goforum26/internal/configs"
	"github.com/panjiasmoroart/goforum26/internal/handlers/memberships"
	"github.com/panjiasmoroart/goforum26/internal/handlers/posts"
	membershipRepo "github.com/panjiasmoroart/goforum26/internal/repository/memberships"
	postRepo "github.com/panjiasmoroart/goforum26/internal/repository/posts"
	membershipSvc "github.com/panjiasmoroart/goforum26/internal/service/memberships"
	postSvc "github.com/panjiasmoroart/goforum26/internal/service/posts"
	"github.com/panjiasmoroart/goforum26/pkg/internalsql"
)

func main() {
	router := gin.Default()

	var (
		cfg *configs.Config
	)

	// Override folder
	// Mengganti folder default dari ./configs menjadi ./internal/configs/.
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()
	log.Println("config : ", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal koneksi ke database", err)
	}

	// testing
	// _ = membershipRepo.NewRepository(db)
	// membershipsHandler := memberships.NewHandler(router)
	// membershipsHandler.RegisterRoute()

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(router, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(router, postService)
	postHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
