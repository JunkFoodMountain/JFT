package repository

type Config struct {
	Retries  int    `envconfig:"default=20,RETRIES"`
	Host     string `envconfig:"default=localhost,PG_HOST"`
	Username string `envconfig:"default=postgres,PG_USERNAME"`
	Password string `envconfig:"default=postgres,PG_PASSWORD"`
}
