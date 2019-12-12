package infra

// Config - config
type Config struct {
	Name       string           `yaml:"name"`
	Mode       string           `yaml:"mode"`
	Server     ServerConfig     `yaml:"server"`
	Log        LogConfig        `yaml:"log"`
	Database   DatabaseConfig   `yaml:"database"`
	Serializer SerializerConfig `yaml:"serializer"`
	Contact    ContactConfig    `yaml:"contact"`
}

// ServerConfig -
type ServerConfig struct {
	Host string `default:"localhost"`
	Port string `default:"3000"`
}

// LogConfig -
type LogConfig struct {
	Code     string `default:"zap"`
	Level    string `default:"info"`
	FileName string `default:"team_action.log" yaml:"file_name"`
}

// DatabaseConfig -
type DatabaseConfig struct {
	Code       string `default:"sqldb"`
	DriverName string `default:"sqlite3" yaml:"driver_name"`
	URLAddress string `yaml:"url_address"`
}

// SerializerConfig -
type SerializerConfig struct {
	Code string `default:"json"`
}

// ContactConfig -
type ContactConfig struct {
	Name  string `default:"Robin Shi"`
	Email string `default:"robinshi@outlook.com"`
}
