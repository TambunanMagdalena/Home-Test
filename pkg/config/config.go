package config

import (
	"net/url"
	"take-home-test/pkg/database"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceHost        string           `mapstructure:"service_host" json:"service_host"`
	ServiceEndpointV   string           `mapstructure:"service_endpoint_v" json:"service_endpoint_v"`
	ServiceEnvironment string           `mapstructure:"service_environment" json:"service_environment"`
	ServicePort        string           `mapstructure:"service_port" json:"service_port"`
	Database           DatabasePlatform `mapstructure:"database" json:"database"`
	JWTSecret          string           `mapstructure:"jwt_secret" json:"jwt_secret"`
	MidtransServerKey  string           `mapstructure:"midtrans_server_key" json:"midtrans_server_key"`
	MidtransClientKey  string           `mapstructure:"midtrans_client_key" json:"midtrans_client_key"`
}

func NewConfig() *Config {
	return &Config{
		ServiceHost:        viper.GetString("APP_HOST"),
		ServiceEndpointV:   viper.GetString("APP_ENDPOINT_V"),
		ServiceEnvironment: viper.GetString("APP_ENVIRONMENT"),
		ServicePort:        viper.GetString("APP_PORT"),
		Database:           LoadDatabaseConfig(),
		JWTSecret:          viper.GetString("JWT_SECRET"),
		MidtransServerKey:  viper.GetString("MIDTRANS_SERVER_KEY"),
		MidtransClientKey:  viper.GetString("MIDTRANS_CLIENT_KEY"),
	}
}

func (c *Config) GetJWTSecret() string {
	if c.JWTSecret == "" {
		return "sagara-tech-secret-key-2025"
	}
	return c.JWTSecret
}

func (d *Database) ToArgs(dbType database.DBType, connType database.ConnType, val url.Values) (res *database.Args) {
	res = &database.Args{
		Username:        d.Username,
		Password:        d.Password,
		Host:            d.URL,
		Port:            d.Port,
		Database:        d.Name,
		Schema:          d.Schema,
		MaxIdleConns:    d.MaxIdleConns,
		MaxOpenConns:    d.MaxOpenConns,
		ConnMaxLifetime: d.MaxLifetime,
		Flavor:          d.Flavor,
		Location:        d.Location,
		Timeout:         d.Timeout,

		DBType:   dbType,
		ConnType: connType,
		Values:   val,
	}
	return
}
