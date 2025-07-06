package entities

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
	Contato
	Endereco
}

type Contato struct {
	Email    string `gorm:"not null" json:"email"`
	Telefone string `gorm:"not null" json:"telefone"`
}

type Endereco struct {
	Cep         string `gorm:"not null" json:"cep"`
	Logradouro  string `gorm:"not null" json:"logradouro"`
	Numero      uint   `gorm:"not null" json:"numero"`
	Complemento string `json:"complemento"`
	Bairro      string `gorm:"not null" json:"bairro"`
	Cidade      string `gorm:"not null" json:"cidade"`
	Estado      string `gorm:"not null" json:"estado"`
}
