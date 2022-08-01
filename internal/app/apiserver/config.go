package apiserver

type Config struct {
	Addr        string `toml:"addr"`
	DatabaseUrl string `toml:"database_url"`
}
