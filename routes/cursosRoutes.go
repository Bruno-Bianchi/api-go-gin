package routes

import (
	"github.com/Bruno-Bianchi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func CursosRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CriaCursos)
	r.GET("/list", controllers.ExibeCursos)
	r.PUT("/update/:id", controllers.EditaCurso)
	r.DELETE("/delete/:id", controllers.DeletaCurso)
	r.GET("/get/:id", controllers.ExibeCursoPorId)
}
