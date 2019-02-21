package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ishansd94/kube-secrets/internal/app/secret"
	"github.com/ishansd94/kube-secrets/pkg/env"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("", secret.Get)
	r.POST("", secret.Create)

	port := env.Get("PORT", "8000")

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}
