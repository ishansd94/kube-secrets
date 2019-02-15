package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ishansd94/kube-secrets/internal/app/secret"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("", secret.Create)

	if err := r.Run(":8000"); err != nil{
		panic(err)
	}
}
