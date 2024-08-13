package routes

import (
	"github.com/Bruno-Bianchi/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AlunosRoutes(r *gin.RouterGroup) {
	r.POST("/", controllers.CriaAlunos)
	r.GET("/list", controllers.ExibeAlunos)
	r.PUT("/update/:id", controllers.EditaAluno)
	r.DELETE("/delete/:id", controllers.DeletaAluno)
	r.GET("/get/:id", controllers.ExibeAlunoPorId)
	r.GET("/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/:nome", controllers.Saudacao)
}
