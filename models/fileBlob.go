package models

type FileBlob struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Mimetype  string `json:"mimetype"`
	Extension string `json:"extension"`
	Type      string `json:"type"` // image, text, video, audio, application
	Content   []byte `json:"content"`
	CreatedAt string `json:"createdAt"`
}
