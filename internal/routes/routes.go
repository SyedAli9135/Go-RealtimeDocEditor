package routes

import (
	"realtime-doc-editor-backend/internal/handlers"
	"realtime-doc-editor-backend/internal/middleware"
	"realtime-doc-editor-backend/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes sets up all application routes
func RegisterRoutes(router *gin.Engine, db *gorm.DB, cfg interface{}) {
	// Initialize the repository
	documentRepo := repositories.NewDocumentRepository(db)

	// Middleware
	router.Use(middleware.LoggerMiddleware())

	// Websocket route
	router.GET("/ws", handlers.WebSocketHandler)

	// Document management routes
	api := router.Group("/api")
	{
		api.POST("/documents", func(c *gin.Context) {
			handlers.CreateDocumentHandler(c, documentRepo)
		})
		api.GET("/documents/:id", func(c *gin.Context) {
			handlers.GetDocumentByIDHandler(c, documentRepo)
		})
		api.PUT("/documents/:id", func(c *gin.Context) {
			handlers.UpdateDocumentHandler(c, documentRepo)
		})
		api.DELETE("/documents/:id", func(c *gin.Context) {
			handlers.DeleteDocumentHandler(c, documentRepo)
		})
	}
}
