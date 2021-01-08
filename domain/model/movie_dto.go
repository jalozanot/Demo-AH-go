package model

import (
	"github.com/jalozanot/demoCeiba/domain/validators"
)

const (
	StatusActive = "active"
)

type MovieDto struct {
	Id           int64  `json:"Id"`
	Nombre       string `json:"Nombre"`
	Categoria    string `json:"Categoria"`
	CodigoBarras string `json:"CodigoBarras"`
}

func (movieDto *MovieDto) CreateMovil(nombre string, categoria string, codigoBarras string) (MovieDto, error) {

	if err := validators.ValidateRequired(nombre, "Nombre should have some value"); err != nil {
		return MovieDto{}, err
	}

	if err := validators.ValidateRequired(categoria, "Categoria should have some value"); err != nil {
		return MovieDto{}, err
	}

	if err := validators.ValidateRequired(categoria, "codigoBarras should have some value"); err != nil {
		return MovieDto{}, err
	}

	return MovieDto{
		Nombre:       nombre,
		Categoria:    categoria,
		CodigoBarras: codigoBarras,
	}, nil
}
