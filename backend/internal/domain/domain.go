package domain

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/hueter57/catalyst/backend/internal/ent"
)

func getenvOrDefault(key, fallback string) string {
	// TODO: これは testutil.GetEnvOrDefault のコピペ
	//       適切な場所に動かしたい
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func loadDsn() string {
	c := mysql.NewConfig()

	c.User = getenvOrDefault("NS_MARIADB_USER", "root")
	c.Passwd = getenvOrDefault("NS_MARIADB_PASSWORD", "pass")
	c.Net = "tcp"
	c.Addr = fmt.Sprintf(
		"%s:%s",
		getenvOrDefault("NS_MARIADB_HOSTNAME", "localhost"),
		getenvOrDefault("NS_MARIADB_PORT", "3306"),
	)
	c.DBName = getenvOrDefault("NS_MARIADB_DATABASE", "app")
	c.Collation = "utf8mb4_general_ci"
	c.AllowNativePasswords = true
	return c.FormatDSN()
}

func Connect() (*ent.Client, error) {
	dsn := loadDsn()
	entOptions := []ent.Option{}
	if os.Getenv("IS_DEBUG_MODE") != "" {
		entOptions = append(entOptions, ent.Debug())
	}
	client, err := ent.Open("mysql", dsn, entOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return client, nil
}
