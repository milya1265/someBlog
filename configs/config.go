package configs

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	Serv struct {
		Port string `yaml:"port"`
	} `yaml:"serv"`

	Storage struct {
		DbDriver string `yaml:"dbDriver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"storage"`
	JWTKey []byte `yaml:"JWTKey"`
}

//     postgres://dmilya:qwerty@localhost:5432/BlogDB?sslmode=disable

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		//D:/Development/Projects/someBlog/config.yaml
		if err := cleanenv.ReadConfig("D:/Development/Projects/someBlog/config.yaml", instance); err != nil {
			log.Fatalln("ERROR --> ", err)
		}
	})
	return instance
}
