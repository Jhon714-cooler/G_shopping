package config

var Conf = new(config)

type config struct {
	System   *ServerConfig `mapstructure:"system" json:"system"`
	Logs     *LogsConfig   `mapstructure:"logs" json:"logs"`
	Database *Database     `mapstructure:"database" json:"database"`
	Mysql    *MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Redis    *RedisConfig  `mapstructure:"redis" json:"redis"`
}

// 定义 ServerConfig 结构体
type ServerConfig struct {
	Address string
	Port    int
}
type Database struct {
	Driver string `mapstructure:"driver" json:"driver"`
	Source string `mapstructure:"source" json:"source"`
}

type MysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type RedisConfig struct {
	// 这里可以添加 Redis 配置的具体字段，例如地址、端口、密码等
	Address  string
	Port     int
	Password string
}

// 定义 LogsConfig 结构体
type LogsConfig struct {
	Level string
	Path  string
}
