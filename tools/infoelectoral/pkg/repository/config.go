package repository

import "fmt"

type Config struct {
	Host     string
	Database string
	User     string
	Password string
}

func (rc Config) toString() string {
	var userPassword string
	host := rc.Host

	if rc.User != "" {
		userPassword = rc.User
		host = "@" + host
		if rc.Password != "" {
			userPassword = userPassword + ":" + rc.Password
		}
	}

	return fmt.Sprintf("%s%s/%s", userPassword, host, rc.Database)
}
