package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Redis  RedisSetting  `mapstructure:"redis"`
	Logger LoggerSetting `mapstructure:"logger"`
	JWT    JWTSetting    `mapstructure:"jwt"`
	SMTP   SMTPSetting   `mapstructure:"smtp"`
}

type ServerSetting struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type MysqlSetting struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Pass         string `mapstructure:"pass"`
	DB           string `mapstructure:"db"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	ConnMaxLife  int    `mapstructure:"conn_max_life"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggerSetting struct {
	LogLevel   string `mapstructure:"log_level"`
	LogFile    string `mapstructure:"log_file"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type JWTSetting struct {
	ExpirationTimeAccessToken  int `mapstructure:"exp_access_token"`
	ExpirationTimeRefreshToken int `mapstructure:"exp_refresh_token"`
}

type SMTPSetting struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}
