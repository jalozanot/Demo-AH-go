package ports

import "github.com/jalozanot/demoCeiba/domain/model"

type MoviesRepository interface {
	Save(user *model.MovieDto) (model.MovieDto, error)
	Get(userId int64) (model.MovieDto, error)
	Update(userId int64, user model.MovieDto) (*model.MovieDto, error)
	Delete(userId int64) error
}
