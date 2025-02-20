package models

type File struct {
	Id        string `json:"name"`
	BucketId  string `json:"bucketId"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	Type      string `json:"type"` // image, text, video, audio, application
	CreatedAt string `json:"createdAt"`
}
