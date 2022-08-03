package sqlstore_test

import (
	"os"
	"testing"
)

var DatabaseURL string

func TestMain(m *testing.M) {
	DatabaseURL = "host=localhost dbname=usertag_test password=123 sslmode=disable"
	os.Exit(m.Run())
}
