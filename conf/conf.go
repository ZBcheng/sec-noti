package conf

// Config ： 配置信息
type Config struct {
	RdConf redisConfig `toml:"redis"`
	PgConf pgConfig    `toml:"postgres"`
}

// RedisConfig : redis配置信息
type redisConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

// PgConfig : Postgres配置信息
type pgConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"username"`
	DBName   string `toml:"dbname"`
	Password string `toml:"password"`
}

// GetConfPath : 获取配置文件路径
func GetConfPath() string {
	return "conf/aliyun_conf.toml"
}