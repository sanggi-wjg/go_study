package server

import "github.com/gin-gonic/gin"

func ServerStart() {
	router := gin.New()
	router.Use(Logger(), Recovery())

	config := NewConfig()
	settings := config.GetConfigSettings()

	router.Static(settings.StaticURL, settings.StaticRoot)
	router.Static(settings.MediaURL, settings.MediaRoot)
	router.StaticFile(settings.FavIconURL, settings.FavIconPath)
	// router.LoadHTMLGlob(settings.TemplateRoot)

	registerRoutes(router, config)

	router.Run(settings.Port)
}
