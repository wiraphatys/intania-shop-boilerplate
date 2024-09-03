package config

type Config interface {
	GetServer() Server
	GetDb() Db
	GetJwt() Jwt
	GetAws() Aws
}

type Server struct {
	Name string `mapstructure:"server_name"`
	Env  string `mapstructure:"server_env"`
	Url  string `mapstructure:"server_url"`
	Host string `mapstructure:"server_host"`
	Port int    `mapstructure:"server_port"`
}

type Db struct {
	Host     string `mapstructure:"db_host"`
	Port     int    `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_pass"`
	Name     string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"db_ssl_mode"`
	Timezone string `mapstructure:"db_timezone"`
}

type Jwt struct {
	ApiSecretKey           string `mapstructure:"jwt_api_secret_key"`
	AccessTokenSecret      string `mapstructure:"jwt_access_token_secret"`
	RefreshTokenSecret     string `mapstructure:"jwt_refresh_token_secret"`
	AccessTokenExpiration  int    `mapstructure:"jwt_access_token_expiration"`
	RefreshTokenExpiration int    `mapstructure:"jwt_refresh_token_expiration"`
}

type Aws struct {
	AwsAccessKeyId     string `mapstructure:"aws_access_key_id"`
	AwsSecretAccessKey string `mapstructure:"aws_secret_access_key"`
	AwsRegion          string `mapstructure:"aws_region"`
}
