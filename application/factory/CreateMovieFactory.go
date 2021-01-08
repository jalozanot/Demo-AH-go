package factory

import (
	"github.com/jalozanot/demoCeiba/application/commands"
	"github.com/jalozanot/demoCeiba/domain/model"
)

func CreateMovie(movieCommand commands.MovieCommand) (model.MovieDto, error) {
	var movieDto model.MovieDto
	movieDto, err := movieDto.CreateMovil(movieCommand.Nombre, movieCommand.Categoria, movieCommand.CodigoBarras)
	return movieDto, err
}
