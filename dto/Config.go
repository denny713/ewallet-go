package dto

type Config struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
}
