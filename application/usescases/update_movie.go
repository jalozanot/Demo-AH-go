package usescases

import (
	"github.com/jalozanot/demoCeiba/application/commands"
	"github.com/jalozanot/demoCeiba/application/factory"
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/domain/ports"
)

type UpdateMovieUseCase interface {
	Handler(userId int64, userCommand commands.MovieCommand) (*model.MovieDto, error)
}

type UseCaseUpdateMovie struct {
	UserRepository ports.MoviesRepository
}

func (useCaseUpdateUser *UseCaseUpdateMovie) Handler(id int64, userCommand commands.MovieCommand) (*model.MovieDto, error) {
	movieDto, err := factory.CreateMovie(userCommand)
	if err != nil {
		return nil, err
	}
	userUpdated, err := useCaseUpdateUser.UserRepository.Update(id, movieDto)
	if err != nil {
		return userUpdated, err
	}
	return userUpdated, nil
}
