package postgres

import "fmt"

func NewConnectionString(user, password, host string, port int, database string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		user,
		password,
		host,
		port,
		database,
	)
}
