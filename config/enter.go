package config

type Config struct{
	Mysql Mysql `yaml:"logger"`	
	Logger Logger `yaml:"mysql"`
	System System `yaml:"system"`
}



