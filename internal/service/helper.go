package service

import (
	mongostore "github.com/junkd0g/go-api-example/internal/persistence/mongo"
	transportkafka "github.com/junkd0g/go-api-example/internal/transport/kafka"
)

func fromKafkaOjectToMongoObject(movie transportkafka.Movie) mongostore.Movie {
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

func fromMongoObjectToKafkaObject(movie mongostore.Movie) *transportkafka.Movie {
	return &transportkafka.Movie{
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
