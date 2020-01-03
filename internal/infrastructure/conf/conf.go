package conf

import "strconv"

type (
	Config interface {
		IsProd() bool
		GetPort() int
		GetDatabaseConfig() DatabaseConfig
		GetRedisConfig() RedisConfig
	}

	ConfigStruct struct {
		ENV      string         `mapstructure:"env"`
		Host     string         `mapstructure:"host"`
		Port     string         `mapstructure:"port"`
		Origins  []OriginConfig `mapstructure:"origins"`
		Database DatabaseConfig `mapstructure:"database"`
		Redis    RedisConfig    `mapstructure:"redis"`
	}

	OriginConfig struct {
		URL string `mapstructure:"url"`
	}

	DatabaseConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Protocol string `mapstructure:"protocol"`
	}

	RedisConfig struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Password string `mapstructure:"password"`
	}
)

func (c *ConfigStruct) IsProd() bool {
	return c.ENV == "prod"
}

func (c *ConfigStruct) IsLocal() bool {
	return c.ENV == "local"
}

func (c *ConfigStruct) GetPort() int {
	if c.Port == "" {
		return 1313
	}

	p, err := strconv.Atoi(c.Port)
	if err != nil {
		return 1313
	}

	return p
}

func (c *ConfigStruct) GetDatabaseConfig() DatabaseConfig {
	if c.Database.Host == "" {
		c.Database.Port = "localhost"
	}

	if c.Database.Port == "" {
		c.Database.Port = "3306"
	}

	if c.Database.User == "" {
		c.Database.Port = "root"
	}

	if c.Database.Protocol == "" {
		c.Database.Protocol = "tcp"
	}

	return c.Database
}

func (c *ConfigStruct) GetRedisConfig() RedisConfig {
	if c.Redis.Port == "" {
		c.Redis.Port = "6379"
	}
	return c.Redis
}
