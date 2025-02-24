package models

type Visibility string
type FileType string

const (
	Public  Visibility = "public"
	Private Visibility = "private"
)

const (
	Image FileType = "image"
	Text  FileType = "text"
)
