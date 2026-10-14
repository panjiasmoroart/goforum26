package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/panjiasmoroart/goforum26/internal/configs"
	"github.com/panjiasmoroart/goforum26/internal/handler/memberships"
	membershipRepo "github.com/panjiasmoroart/goforum26/internal/repository/memberships"
	membershipSvc "github.com/panjiasmoroart/goforum26/internal/service/memberships"
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

	membershipService := membershipSvc.NewService(membershipRepo)

	membershipHandler := memberships.NewHandler(router, membershipService)
	membershipHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
