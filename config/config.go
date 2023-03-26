package config

type Config struct {
	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string

	DefaultOffset int
	DefaultLimit  int
}

func Load() Config {
	cfg := Config{}

	cfg.ServerHost = "localhost"
	cfg.ServerPort = ":4001"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "khumoyun"
	cfg.PostgresDatabase = "shop_product"
	cfg.PostgresPassword = "admin@111"
	cfg.PostgresPort = "5432"
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	return cfg
}
