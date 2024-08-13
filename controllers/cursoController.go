package controllers

import (
	"net/http"

	"github.com/Bruno-Bianchi/api-go-gin/database"
	"github.com/Bruno-Bianchi/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func CriaCursos(c *gin.Context) {
	var curso models.Curso
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&curso)
	c.JSON(http.StatusOK, curso)
}

func ExibeCursos(c *gin.Context) {
	var cursos []models.Curso
	database.DB.Find(&cursos)
	c.JSON(200, cursos)
}

func EditaCurso(c *gin.Context) {
	var curso models.Curso
	id := c.Params.ByName("id")
	database.DB.First(&curso, id)

	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Model(&curso).UpdateColumns(curso)
	c.JSON(http.StatusOK, curso)
}

func DeletaCurso(c *gin.Context) {
	var curso models.Curso
	id := c.Params.ByName("id")
	database.DB.Delete(&curso, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Curso deletado com sucesso!",
	})
}

func ExibeCursoPorId(c *gin.Context) {
	var curso models.Curso
	id := c.Params.ByName("id")
	database.DB.First(&curso, id)

	if curso.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Curso não encontrado!",
		})
		return
	}
	c.JSON(http.StatusOK, &curso)
}
