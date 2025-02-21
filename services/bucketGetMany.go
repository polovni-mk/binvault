package services

import (
	"binvault/database"
	"binvault/models"
)

func BucketGetMany(limit int, offset int) []models.Bucket {
	db := database.ObtainConnection()
	var buckets []models.Bucket
	var entries []database.Bucket
	db.Limit(limit).Offset(offset).Find(&entries)
	for _, entry := range entries {
		buckets = append(buckets, models.Bucket{
			Name:       entry.Name,
			CreatedBy:  &entry.CreatedBy,
			CreatedAt:  entry.CreatedAt,
			Visibility: entry.Visibility,
		})
	}
	return buckets

}
