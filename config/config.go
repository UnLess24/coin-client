package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB             DB
	Server         Server
	MigrationsPath string
	JWTSecretKey   string
	CoinServer     CoinServerConnect
}

type DB struct {
	Name     string
	Host     string
	Port     string
	Database string
	SslMode  string
	User     string
	Password string
}
type Server struct {
	Host        string
	Port        string
	ReleaseMode bool
}

type CoinServerConnect struct {
	Schema string
	Host   string
	Port   string
	Type   string
}

func MustRead() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return &Config{
		Server: Server{
			Host:        viper.GetString("server.host"),
			Port:        viper.GetString("server.port"),
			ReleaseMode: viper.GetBool("server.release_mode"),
		},
		DB: DB{
			Name:     viper.GetString("db.name"),
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Database: viper.GetString("db.database"),
			SslMode:  viper.GetString("db.sslmode"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
		},
		MigrationsPath: viper.GetString("migrations.path"),
		JWTSecretKey:   viper.GetString("jwt.secret"),
		CoinServer: CoinServerConnect{
			Schema: viper.GetString("coin-server-connect.schema"),
			Host:   viper.GetString("coin-server-connect.host"),
			Port:   viper.GetString("coin-server-connect.port"),
			Type:   viper.GetString("coin-server-connect.type"),
		},
	}
}
