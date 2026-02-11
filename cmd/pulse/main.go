package main

import (
	"os"
	
	"github.com/eliferdentr/pulse/internal/api"
	j "github.com/eliferdentr/pulse/internal/jobs"
	"github.com/eliferdentr/pulse/internal/logger"
)

func main() {
	// r := gin.Default()
	logger.Init()
	store := j.NewStore()
	manager := j.NewManager(store, 13)
	manager.StartWorkers(3)
	// Router'ı oluştur
	r := api.NewRouter(manager)

	// Port ayarla (ENV > default)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Server başlat
	if err := r.Run(":" + port); err != nil {
		logger.Log.Error("failed to start server", "error", err)
	}
}
