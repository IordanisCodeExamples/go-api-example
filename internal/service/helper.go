package service

import (
	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	"github.com/junkd0g/go-api-example/internal/transport"
)

func fromKafkaOjectToMongoObject(movie transport.Movie) mongostore.Movie {
	return mongostore.Movie{
		Title:            movie.Title,
		Year:             movie.Year,
		Duration:         movie.Duration,
		Director:         movie.Director,
		Cast:             movie.Cast,
		Genre:            movie.Genre,
		Synopsis:         movie.Synopsis,
		BoxOfficeRevenue: movie.BoxOfficeRevenue,
	}
}

func fromMongoObjectToKafkaObject(movie mongostore.Movie) *transport.Movie {
	return &transport.Movie{
		Title:            movie.Title,
		Year:             movie.Year,
		Duration:         movie.Duration,
		Director:         movie.Director,
		Cast:             movie.Cast,
		Genre:            movie.Genre,
		Synopsis:         movie.Synopsis,
		BoxOfficeRevenue: movie.BoxOfficeRevenue,
	}
}
