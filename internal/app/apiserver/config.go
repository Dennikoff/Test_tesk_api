package apiserver

type Config struct {
	Addr        string `toml:"addr"`
	DatabaseURL string `toml:"database_url"`
	DriverName  string `toml:"driver_name"`
}
