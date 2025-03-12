package interfaces

type Database struct {
	Host     string
	User     string
	Database string
	Port     string
	Password string
}
type AppConfig struct {
	Host string
	Port string
}
type Storage struct {
	UploadPath string
	BackUpPath string
}
type Config struct {
	AppConfig AppConfig
	Storage   Storage
	Database  Database
}
