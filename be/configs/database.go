package configs

import "github.com/badrunp/media-sosial-web/be/helpers"

type database struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	SslMode string
}

func DatabaseConfig() *database {
	return &database{
		Name:     helpers.GetEnvStr("DB_NAME", ""),
		Host:     helpers.GetEnvStr("DB_HOST", ""),
		Port:     helpers.GetEnvNum("DB_PORT", 8080),
		User:     helpers.GetEnvStr("DB_USER", ""),
		Password: helpers.GetEnvStr("DB_PASS", ""),
		SslMode: helpers.GetEnvStr("SSL_MODE", ""),
	}
}

func TestDatabaseConfig() *database {
	return &database{
		Name:     helpers.GetEnvStr("DB_TEST_NAME", ""),
		Host:     helpers.GetEnvStr("DB_TEST_HOST", ""),
		Port:     helpers.GetEnvNum("DB_TEST_PORT", 8080),
		User:     helpers.GetEnvStr("DB_TEST_USER", ""),
		Password: helpers.GetEnvStr("DB_TEST_PASS", ""),
		SslMode: helpers.GetEnvStr("SSL_TEST_MODE", ""),
	}
}
