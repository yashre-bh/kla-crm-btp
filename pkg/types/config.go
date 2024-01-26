package types

type Config struct {
	Database DatabaseConfig `toml:"database"`
	JWT      JWTConfig      `toml:"jwt"`
}

type DatabaseConfig struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Name     string `toml:"name"`
}

type JWTConfig struct {
	Secret string `toml:"secret"`
}
