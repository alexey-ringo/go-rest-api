package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

//Вызывается один раз перед всеми тестами в пакете
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test user=default password=secret sslmode=disable"
	}

	os.Exit(m.Run())

}
