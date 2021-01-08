package movie_rep

import (
	"errors"
	"fmt"

	"github.com/fmcarrero/bookstore_utils-go/logger"
	"github.com/jalozanot/demoCeiba/domain/exceptions"
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/models"
	"github.com/jalozanot/demoCeiba/infrastructure/database_client"
	"github.com/jalozanot/demoCeiba/infrastructure/mappers/users_mapper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserMysqlRepository struct {
	Db *gorm.DB
}

const (
	SET = "SET"
	GET = "GET"
)

func (userMysqlRepository *UserMysqlRepository) Save(user *model.MovieDto) (model.MovieDto, error) {

	var movieEntity models.MovieEntity
	movieEntity = users_mapper.MovieDtoToMovieEntity(*user)
	if err := userMysqlRepository.Db.Create(&movieEntity).Error; err != nil {
		logger.Error(fmt.Sprintf("can't work with %s", movieEntity.Nombre), err)
		return model.MovieDto{}, errors.New(fmt.Sprintf("can't work with %s", movieEntity.Categoria))
	}
	userUpdated := users_mapper.MovieEntityToMovieDto(movieEntity)
	status, _ := database_client.Con.Do(SET, movieEntity.CodigoBarras, movieEntity)
	fmt.Println(status)
	return userUpdated, nil
}

func (userMysqlRepository *UserMysqlRepository) Get(userId int64) (model.MovieDto, error) {

	var movieEntity models.MovieEntity
	if userMysqlRepository.Db.First(&movieEntity, userId).Error != nil {

		return model.MovieDto{}, exceptions.MovieNotFound{ErrMessage: fmt.Sprintf("movie with id=%d not found", userId)}
	}

	movieDto := users_mapper.MovieEntityToMovieDto(movieEntity)
	return movieDto, nil
}

func (userMysqlRepository *UserMysqlRepository) Update(userId int64, movieDto model.MovieDto) (*model.MovieDto, error) {
	var current models.MovieEntity
	if userMysqlRepository.Db.First(&current, userId).RecordNotFound() {
		return nil, errors.New(fmt.Sprintf("movie not found %v", userId))
	}
	if userMysqlRepository.Db.Model(&current).Update(models.MovieEntity{Nombre: movieDto.Nombre, Categoria: movieDto.Categoria, CodigoBarras: movieDto.CodigoBarras}).Error != nil {
		return nil, errors.New(fmt.Sprintf("error when updated user %v", userId))
	}
	userUpdated := users_mapper.MovieEntityToMovieDto(current)
	return &userUpdated, nil
}

func (userMysqlRepository *UserMysqlRepository) Delete(Id int64) error {

	// var current models.MovieEntity
	// current.ID = Id
	if userMysqlRepository.Db.Delete(Id).Error != nil {
		return errors.New(fmt.Sprintf("cannot delete movie  %v", Id))
	}
	return nil
}
