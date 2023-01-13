package controllers

import (
	"net/http"
	"github.com/go/crud/models"
	"github.com/go/crud/services"
	"github.com/gin-gonic/gin"
)

type MovieController struct {
	MovieService services.MovieService
}

func New(Movieservice services.MovieService) MovieController {
	return MovieController{
		MovieService: movieservice,
	}
}

func (uc *MovieController) CreateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.MovieService.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *MovieController) GetMovie(ctx *gin.Context) {
	var moviename string = ctx.Param("name")
	movie, err := uc.MovieService.GetMovie(&moviename)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (uc *MovieController) GetAll(ctx *gin.Context) {
	movies, err := uc.MovieService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (uc *MovieController) CommentOnMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.MovieService.CommentOnMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *MovieController) LikeTheMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.MovieService.LikeTheMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
func (uc *MovieController) UpdateMovie(ctx *gin.Context) {
	var movie models.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.MovieService.UpdateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *MovieController) DeleteMovie(ctx *gin.Context) {
	var moviename string = ctx.Param("name")
	err := uc.MovieService.DeleteMovie(&moviename)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *MovieController) RegisterMovieRoutes(rg *gin.RouterGroup) {
	movieroute := rg.Group("/movie")
	movieroute.POST("/create", uc.CreateMovie)
	movieroute.GET("/get/:name", uc.GetMovie)
	movieroute.GET("/getall", uc.GetAll)
	movieroute.PATCH("/:id/like", uc.LikeTheMovie)
	movieroute.PATCH("/:id/comment", uc.CommentOnMovie)
	movieroute.PATCH("/update/:id", uc.UpdateMovie)
	movieroute.DELETE("/delete/:id", uc.DeleteMovie)
}
