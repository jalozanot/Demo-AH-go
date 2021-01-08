package usescases

import (
	"github.com/jalozanot/demoCeiba/application/commands"
	"github.com/jalozanot/demoCeiba/application/factory"
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/domain/ports"
)

type CreatesMoviePort interface {
	Handler(movieCommand commands.MovieCommand) (model.MovieDto, error)
}

type UseCaseMovieCreate struct {
	MovieRepository ports.MoviesRepository
}

func (createsUseCase *UseCaseMovieCreate) Handler(movieCommand commands.MovieCommand) (model.MovieDto, error) {

	movieDto, err := factory.CreateMovie(movieCommand)
	if err != nil {
		return model.MovieDto{}, err
	}
	movie, createMovieErr := createsUseCase.MovieRepository.Save(&movieDto)
	if createMovieErr != nil {
		return model.MovieDto{}, createMovieErr
	}
	return movie, nil

}
