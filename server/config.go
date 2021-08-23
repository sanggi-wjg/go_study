package server

type Config struct {
	settings *Settings
	database string
}

func NewConfig() *Config {
	config := Config{
		settings: NewSettings(),
		database: "",
	}

	return &config
}

func (config *Config) GetConfigSettings() *Settings {
	return config.settings
}
