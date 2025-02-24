package services

import (
	"binvault/pkg/database"
	"binvault/pkg/models"
)

func FileGetOne(bucketName string, fileName string) models.File {
	db := database.ObtainConnection()
	var bucket database.Bucket
	db.First(&bucket, "name = ?", bucketName)
	var entry database.File
	db.First(&entry, "name = ? bucketID = ?", fileName, bucket.ID)
	return models.File{
		Bucket:     bucket.Name,
		Name:       entry.Name,
		Size:       entry.Size,
		Extension:  entry.Extension,
		Type:       entry.Type,
		Visibility: bucket.Visibility,
	}
}
