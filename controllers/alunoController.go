package controllers

import (
	"net/http"

	"github.com/Bruno-Bianchi/api-go-gin/database"
	"github.com/Bruno-Bianchi/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

type AlunoInput struct {
	Nome    string `json:"nome" binding:"required"`
	CPF     string `json:"cpf" binding:"required"`
	RG      string `json:"rg" binding:"required"`
	CursoID uint   `json:"curso_id" binding:"required"` // ID do curso
}

func ExibeAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func ExibeAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado!",
		})
		return
	}
	c.JSON(http.StatusOK, &aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaAlunos(c *gin.Context) {
	var input AlunoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cria uma instância do modelo Aluno
	aluno := models.Aluno{
		Nome: input.Nome,
		CPF:  input.CPF,
		RG:   input.RG,
	}

	// Salva o aluno no banco de dados
	if err := database.DB.Create(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Associa o curso ao aluno
	var curso models.Curso
	if err := database.DB.First(&curso, input.CursoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curso não encontrado"})
		return
	}

	// Associa o curso ao aluno
	if err := database.DB.Model(&aluno).Association("Cursos").Append(&curso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso!",
	})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado!",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}
