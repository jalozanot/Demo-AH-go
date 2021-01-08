package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/application/usescases"
	"github.com/jalozanot/demoCeiba/domain/ports"
	"github.com/jalozanot/demoCeiba/infrastructure/adapters/repository/movie_rep"
	"github.com/jalozanot/demoCeiba/infrastructure/app/middlewares/error_handler"
	"github.com/jalozanot/demoCeiba/infrastructure/controllers"
	"github.com/jalozanot/demoCeiba/infrastructure/database_client"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

//StartApplication inicio de aplicacion
func StartApplication() {

	_ = godotenv.Load()
	router.Use(error_handler.ErrorHandler())
	userRepository := getUsersRepository()
	var handler = createHandler(userRepository)
	r := mapUrls(handler)

	database_client.GetConnectionRedis()

	_ = r.Run(":8084")
}

func createHandler(userRepository ports.MoviesRepository) controllers.RedirectMovieHandler {

	return newHandler(newCreatesUseCase(userRepository), newGetMovieUseCase(userRepository),
		newUpdateMovieUseCase(userRepository), newDeleteMovieUseCase(userRepository))
}
func newCreatesUseCase(repository ports.MoviesRepository) usescases.CreatesMoviePort {
	return &usescases.UseCaseMovieCreate{
		MovieRepository: repository,
	}
}

func newGetMovieUseCase(repository ports.MoviesRepository) usescases.GetMovieUseCase {
	return &usescases.UseCaseGetMovie{
		UserRepository: repository,
	}
}

func newUpdateMovieUseCase(repository ports.MoviesRepository) usescases.UpdateMovieUseCase {
	return &usescases.UseCaseUpdateMovie{
		UserRepository: repository,
	}
}

func newDeleteMovieUseCase(usersRepository ports.MoviesRepository) usescases.DeleteMovieUseCase {
	return &usescases.UseCaseDeleteMovie{
		UserRepository: usersRepository,
	}
}

func newHandler(createMovie usescases.CreatesMoviePort, getMovieUseCase usescases.GetMovieUseCase, updateMovieUseCase usescases.UpdateMovieUseCase,
	deleteMovieUseCase usescases.DeleteMovieUseCase) controllers.RedirectMovieHandler {
	return &controllers.Handler{CreatesUseCase: createMovie, GetMovieUseCase: getMovieUseCase, UseCaseUpdateMovie: updateMovieUseCase,
		UseCaseDeleteMovie: deleteMovieUseCase,
	}
}
func getUsersRepository() ports.MoviesRepository {
	return &movie_rep.UserMysqlRepository{
		Db: database_client.GetDatabaseInstance(),
	}
}
