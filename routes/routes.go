package routes

import (
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(gin.Recovery())

	cursosGroup := r.Group("/cursos")
	CursosRoutes(cursosGroup)

	alunosGroup := r.Group("/alunos")
	AlunosRoutes(alunosGroup)

	r.Run()
}
