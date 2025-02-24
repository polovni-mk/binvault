package services

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func BucketGetOne(bucketName string) models.Bucket {
	db := database.ObtainConnection()
	var entry database.Bucket
	db.First(&entry, "name = ?", bucketName)
	return models.Bucket{
		Name:       entry.Name,
		CreatedBy:  &entry.CreatedBy,
		CreatedAt:  entry.CreatedAt,
		Visibility: entry.Visibility,
	}
}
