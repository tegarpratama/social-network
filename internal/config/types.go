package config

type ConfigTypes struct {
	PORT string `mapstructure:"port"`

	DB_HOST     string `mapstructure:"db_host"`
	DB_PORT     string `mapstructure:"db_port"`
	DB_USER     string `mapstructure:"db_user"`
	DB_PASSWORD string `mapstructure:"db_password"`
	DB_DATABASE string `mapstructure:"db_database"`

	SECRET_JWT string `mapstructure:"secret_jwt"`
}
