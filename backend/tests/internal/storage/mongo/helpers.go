package mongo

import "go.mongodb.org/mongo-driver/bson"

// getProjectionForAnwers return projection for all existing bson answer fields in domain entities
// It used to get quick bson to exclude all answer fields from bson document
func getProjectionForAnwers() bson.D {
	return bson.D{{"questions.answers", 0}}
}
