package users_mapper

import (
	"github.com/jalozanot/demoCeiba/domain/model"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/models"
)

func MovieDtoToMovieEntity(movieDto model.MovieDto) models.MovieEntity {

	var movieEntity models.MovieEntity
	movieEntity.Nombre = movieDto.Nombre
	movieEntity.Categoria = movieDto.Categoria
	movieEntity.CodigoBarras = movieDto.CodigoBarras

	return movieEntity
}

func MovieEntityToMovieDto(movieEntity models.MovieEntity) model.MovieDto {
	var MovieDto model.MovieDto
	MovieDto.Id = movieEntity.ID
	MovieDto.Nombre = movieEntity.Nombre
	MovieDto.Categoria = movieEntity.Categoria
	MovieDto.CodigoBarras = movieEntity.CodigoBarras
	return MovieDto
}

func MovieEntityToMovilesDto(MovieEntity []models.MovieEntity) []model.MovieDto {
	var movies []model.MovieDto
	for _, MovieEntity := range MovieEntity {
		movie := MovieEntityToMovieDto(MovieEntity)
		movies = append(movies, movie)
	}
	return movies
}
