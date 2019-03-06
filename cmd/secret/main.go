package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ishansd94/kube-secrets/internal/app/secret"
	"github.com/ishansd94/kube-secrets/internal/pkg/version"
	"github.com/ishansd94/kube-secrets/pkg/env"
	"github.com/ishansd94/kube-secrets/pkg/log"
)

func main() {

	log.Info("main", fmt.Sprintf("starting service... commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release))

	exit := make(chan struct{})

	srv := Router()

	// Background listener for os events.
	go shutdownHandler(srv, exit)

	log.Info("main", fmt.Sprintf("listening on %s", srv.Addr))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("main", "error starting server...", err)
	}

	<-exit

	log.Info("main", "server exiting...")
}

func Router() *http.Server {

	r := gin.Default()

	r.Use(cors.Default())

	// Application common endpoints
	r.GET("/version", version.Get)

	// Application specific endpoints
	r.GET("", secret.Get)
	r.POST("", secret.Create)

	port := env.Get("PORT", "8000")

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}
}

func shutdownHandler(srv *http.Server, ch chan struct{}) {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGKILL, syscall.SIGTERM)
	ctx := context.Background()

	log.Warn("main", fmt.Sprintf("server shutting down.. got %s", <-interrupt))

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("main", "error shutting down server...", err)
	}

	close(ch)
}
