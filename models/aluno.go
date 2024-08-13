package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome   string  `json:"nome"`
	CPF    string  `json:"cpf"`
	RG     string  `json:"rg"`
	Cursos []Curso `gorm:"many2many:aluno_curso;"`
}
