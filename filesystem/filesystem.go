package filesystem

import (
	"binvault/models"
	"binvault/utils"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func GetBuckets() []*models.Bucket {
	path := utils.GetEnv("FILE_SYS_PATH")
	var dirNames []string = readDirectories(path)
	var buckets []*models.Bucket
	for _, name := range dirNames {
		bucket := GetBucket(name)
		buckets = append(buckets, bucket)
	}
	return buckets
}

func GetBucket(name string) *models.Bucket {
	bucket := createBucketModel(name)
	return bucket
}

func readDirectories(path string) []string {
	var dirs []string
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs
}

func createBucketModel(bucketId string) *models.Bucket {
	path := utils.GetEnv("FILE_SYS_PATH")
	filePath := filepath.Join(path, bucketId, ".meta")
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening metadata file:", err)
		return &models.Bucket{}
	}
	defer file.Close()
	var meta models.Bucket
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&meta)
	if err != nil {
		log.Println("Error decoding metadata file:", err)
		return &models.Bucket{}
	}
	meta.Id = bucketId
	return &meta
}
