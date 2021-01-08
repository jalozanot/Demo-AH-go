package usescases

import (
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/domain/ports"
)

type GetMovieUseCase interface {
	Handler(userId int64) (model.MovieDto, error)
}

type UseCaseGetMovie struct {
	UserRepository ports.MoviesRepository
}

func (useCaseGetMovie *UseCaseGetMovie) Handler(id int64) (model.MovieDto, error) {

	movieDto, err := useCaseGetMovie.UserRepository.Get(id)
	if err != nil {
		return model.MovieDto{}, err
	}
	return movieDto, nil
}
