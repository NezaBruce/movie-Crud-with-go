package services

import (
	"context"
	"errors"

	"github.com/go/crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type MovieServiceImpl struct {
	moviecollection *mongo.Collection
	ctx            context.Context
}

func NewMovieService(moviecollection *mongo.Collection, ctx context.Context) MovieService {
	return &MovieServiceImpl{
		moviecollection: moviecollection,
		ctx:            ctx,
	}
}

func (u *MovieServiceImpl) CreateMovie(movie *models.Movie) error {
	_, err := u.moviecollection.InsertOne(u.ctx, movie)
	return err
}

func (u *MovieServiceImpl) GetMovie(name *string) (*models.Movie, error) {
	var movie *models.Movie
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.moviecollection.FindOne(u.ctx, query).Decode(&movie)
	return movie, err
}

func (u *MovieServiceImpl) GetAll() ([]*models.Movie, error) {
	var movies []*models.Movie
	cursor, err := u.moviecollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var movie models.Movie
		err := cursor.Decode(&movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(movies) == 0 {
		return nil, errors.New("movies not found")
	}
	return movies, nil
}

func (u *MovieServiceImpl) UpdateMovie(movie *models.Movie) error {
	filter := bson.D{primitive.E{Key: "id", Value: movie.id}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: movie.Name}, primitive.E{Key: "banner", Value: movie.Banner}, primitive.E{Key: "cast", Value: movie.Cast}}}}
	result, _ := u.moviecollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched movie found for update")
	}
	return nil
}

func (u *MovieServiceImpl) DeleteMovie(id *int) error {
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	result, _ := u.moviecollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched movie found for delete")
	}
	return nil
}
