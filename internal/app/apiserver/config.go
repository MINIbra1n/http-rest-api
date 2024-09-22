package apiserver

//Config ...

type Config struct {
	BildAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BildAddr: ":8080",
		LogLevel: "debug",
	}
}
