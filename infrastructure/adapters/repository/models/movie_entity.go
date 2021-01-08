package models

type MovieEntity struct {
	ID           int64
	Nombre       string
	Categoria    string
	CodigoBarras string
}

func (MovieEntity) TableName() string {
	return "movie"
}
