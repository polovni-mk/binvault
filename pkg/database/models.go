package database

import (
	"binvault/pkg/models"

	"gorm.io/gorm"
)

type Bucket struct {
	gorm.Model
	Name       string `gorm:"unique"`
	CreatedBy  string
	Visibility models.Visibility
	Files      []File `gorm:"foreignKey:BucketID"`
}
type File struct {
	gorm.Model
	BucketID  uint
	Bucket    Bucket `gorm:"foreignKey:BucketID"`
	Name      string `gorm:"unique"`
	Size      int64
	Extension string
	Type      models.FileType
	Hash      string
}
