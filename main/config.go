package main

type Config struct {
	useDb         bool `mapstructure:"path_map"`
	DBSource      string 
	ServerAddress string
}
