package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/panjiasmoroart/goforum26/internal/configs"
	"github.com/panjiasmoroart/goforum26/internal/handler/memberships"
)

func main() {
	router := gin.Default()

	var (
		cfg *configs.Config
	)

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

	membershipsHandler := memberships.NewHandler(router)
	membershipsHandler.RegisterRoute()

	router.Run(cfg.Service.Port)
}
