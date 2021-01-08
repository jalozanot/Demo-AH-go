package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fmcarrero/bookstore_utils-go/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/application/commands"
	"github.com/jalozanot/demoCeiba/application/usescases"
	"github.com/jalozanot/demoCeiba/infrastructure/marshallers"
)

type RedirectMovieHandler interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type Handler struct {
	CreatesUseCase     usescases.CreatesMoviePort
	GetMovieUseCase    usescases.GetMovieUseCase
	UseCaseUpdateMovie usescases.UpdateMovieUseCase
	UseCaseDeleteMovie usescases.DeleteMovieUseCase
}

func (h *Handler) Create(c *gin.Context) {

	var userCommand commands.MovieCommand
	if err := c.ShouldBindJSON(&userCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	result, createUserErr := h.CreatesUseCase.Handler(userCommand)

	if createUserErr != nil {
		_ = c.Error(createUserErr)
		return
	}

	isPublic := true
	c.JSON(http.StatusCreated, marshallers.Marshall(isPublic, result))
}
func (h *Handler) Get(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	fmt.Println(userId)
	movieDto, errGet := h.GetMovieUseCase.Handler(userId)
	if errGet != nil {
		_ = c.Error(errGet)
		return
	}

	// if oauth.GetCallerId(c.Request) == user.Id {
	// 	c.JSON(http.StatusOK, marshallers.Marshall(false, user))
	// 	return
	// }
	//c.JSON(http.StatusOK, marshallers.Marshall(oauth.IsPublic(c.Request), model.MovieDto{}))
	c.JSON(http.StatusOK, marshallers.Marshall(true, movieDto))
}

func (h *Handler) Update(c *gin.Context) {
	id, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	var movieCommand commands.MovieCommand
	if err := c.ShouldBindJSON(&movieCommand); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json")
		c.JSON(restErr.Status(), restErr)
		return
	}
	movieDto, updateErr := h.UseCaseUpdateMovie.Handler(id, movieCommand)
	if updateErr != nil {
		restErr := rest_errors.NewBadRequestError(updateErr.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}

	c.JSON(http.StatusOK, &movieDto)
}

func (h *Handler) Delete(c *gin.Context) {
	id, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		restErr := rest_errors.NewBadRequestError("id should be valid")
		c.JSON(restErr.Status(), restErr)
		return
	}
	errDelete := h.UseCaseDeleteMovie.Handler(id)
	if errDelete != nil {
		restErr := rest_errors.NewBadRequestError(errDelete.Error())
		c.JSON(restErr.Status(), restErr)
		return
	}
	c.Status(http.StatusNoContent)
}
