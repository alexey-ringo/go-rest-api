package apiserver

type Config struct {
	//Адрес, на котором запускается веб-сервер
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

//Инициализированный конфиг с дефолтными параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8081",
		LogLevel: "debug",
	}
}
