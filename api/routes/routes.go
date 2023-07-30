package routes

import (
	"github.com/Atoo35/gingonic-surrealdb/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/surrealdb/surrealdb.go"
)

func SetupRoutes(db *surrealdb.DB) *gin.Engine {
	h := &handlers.TaskHandler{DB: db}
	router := gin.Default()

	tasksRoutes := router.Group("/api/tasks")
	{
		tasksRoutes.GET("/", h.GetTasks)
		tasksRoutes.POST("/", h.CreateTask)
		tasksRoutes.GET("/:id", h.GetTask)
		tasksRoutes.PUT("/:id", h.UpdateTask)
		tasksRoutes.DELETE("/:id", h.DeleteTask)
	}

	return router
}
