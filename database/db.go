package database

import (
	"log"

	"github.com/Bruno-Bianchi/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=alunos password=123 dbname=alunos port=5432 sslmode=disable"

	// Abre a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Migrações
	// Curso
	if err := DB.AutoMigrate(&models.Curso{}); err != nil {
		log.Fatalf("Erro ao migrar banco de dados: %v", err)
	}

	// Aluno
	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Fatalf("Erro ao migrar banco de dados: %v", err)
	}
}
