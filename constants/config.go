package constants

import (
	"fmt"
	"net/url"
	"strconv"
	"sync"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConstants struct {
	PORT               string `env:"PORT" envDefault:"42069"`
	HOST               string `env:"HOST" `
	DBPort             int    `env:"DB_PORT"`
	DBUser             string `env:"DB_USER"`
	DBPassword         string `env:"DB_PASS"`
	DBName             string `env:"DB_NAME"`
	DBSslMode          string `env:"DB_SSL_MODE" envDefault:"disable"`
	JWTSecret          string `env:"JWT_SECRET"`
	RefreshSecret      string `env:"RERESH_SECRET"`
	AccessTokenMinutes int    `env:"ACCESS_TOKEN_MINUTES" envDefault:"15"`
	RefreshTokenHours  int    `env:"REFRESH_TOKEN_HOURS" envDefault:"168"`
	DatabaseURl        string `env:"DATABASE_URL"`
}

var (
	configOnce   sync.Once
	envConstants EnvConstants
	configErr    error
)

func GetEnv() (EnvConstants, error) {
	configOnce.Do(func() {
		_ = godotenv.Load()

		if err := env.Parse(&envConstants); err != nil {
			configErr = err
			return
		}

		if envConstants.PORT[0] != ':' {
			envConstants.PORT = fmt.Sprintf(":%s", envConstants.PORT)
		}
		if envConstants.DatabaseURl != "" &&
			(envConstants.HOST == "" || envConstants.DBUser == "" || envConstants.DBPassword == "") {

			parsed, err := url.Parse(envConstants.DatabaseURl)
			if err != nil {
				configErr = err
				return
			}

			if parsed.User != nil {
				envConstants.DBUser = parsed.User.Username()
				pass, _ := parsed.User.Password()
				envConstants.DBPassword = pass
			}

			envConstants.HOST = parsed.Hostname()

			portStr := parsed.Port()
			if portStr != "" {
				port, _ := strconv.Atoi(portStr)
				envConstants.DBPort = port
			}

			envConstants.DBName = parsed.Path[1:]

			q := parsed.Query()
			if ssl := q.Get("sslmode"); ssl != "" {
				envConstants.DBSslMode = ssl
			}
		}

		if envConstants.HOST == "" || envConstants.DBUser == "" || envConstants.DBPassword == "" {
			configErr = ErrMissingENV
		}
	})

	return envConstants, configErr
}
