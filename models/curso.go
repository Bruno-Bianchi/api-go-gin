package models

import "gorm.io/gorm"

type Curso struct {
	gorm.Model
	Descricao string `json:"descricao"`
}
