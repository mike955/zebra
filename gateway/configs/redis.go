package configs

type Redis struct {
	RedisAddr          string `yaml:"redis_addr"`
	RedisPassword      string `yaml:"redis_password"`
	RedisDefaultExpire string `yaml:"redis_default_expire"`
}
