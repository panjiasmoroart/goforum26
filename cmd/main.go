package main

import (
	"github.com/gin-gonic/gin"
	"github.com/panjiasmoroart/goforum26/internal/handler/memberships"
)

func main() {
	router := gin.Default()

	membershipsHandler := memberships.NewHandler(router)
	membershipsHandler.RegisterRoute()

	router.Run(":9999")
}
