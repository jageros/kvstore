package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

var CFG *Config

type Config struct {
	DBAddr     string `toml:"dbAddr"`
	DBName     string `toml:"dbName"`
	DBUser     string `toml:"dbUser"`
	DBPassword string `toml:"dbPassword"`
	GinModel   string `toml:"ginModel"`
	ListenAddr string `toml:"listenAddr"`
}

func loadConfig(configPath string) {
	var config *Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Panic(err)
	}
	CFG = config
}

// ======== init ==============
func Initialize(configPath string) {
	loadConfig(configPath)
}
