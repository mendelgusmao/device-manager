package application

type Configuration struct {
	ServerAddress string `default:":8080"`
	DSN           string `default:"device-manager.db"`
}
