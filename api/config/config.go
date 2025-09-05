package config

type Config struct{
	Port int
}

func LoadConfig() (*Config,error){
	var c Config
	c.Port = 8080
	return &c,nil
}