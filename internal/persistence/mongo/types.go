package mongostore

import "go.mongodb.org/mongo-driver/bson/primitive"

// Movie represents the movie model in the database
type Movie struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Title            string             `bson:"title"`
	Year             int                `bson:"year"`
	Duration         int                `bson:"duration"` // In minutes
	Director         string             `bson:"director"`
	Cast             []string           `bson:"cast"`
	Genre            []string           `bson:"genre"`
	Synopsis         string             `bson:"synopsis"`
	BoxOfficeRevenue float64            `bson:"box_office_revenue"`
}
