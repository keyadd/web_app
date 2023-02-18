package config

type Server struct {
	JWT       JWT       `mapstructure:"jwt" json:"jwt" yaml:"jwt"`       //Jwt 配置
	Zap       Zap       `mapstructure:"zap" json:"zap" yaml:"zap"`       //zap 日志配置
	Redis     Redis     `mapstructure:"redis" json:"redis" yaml:"redis"` //redis 配置
	Email     Email     `mapstructure:"email" json:"email" yaml:"email"`
	Casbin    Casbin    `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System    System    `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha   `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Snowflake Snowflake `mapstructure:"snowflake" json:"snowflake" yaml:"snowflake"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"` //mysql 配置
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	Minio Minio `mapstructure:"minio" json:"minio" yaml:"minio"` //minio开源对象存储
}
