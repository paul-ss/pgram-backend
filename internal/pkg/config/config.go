package config

import (
	"github.com/spf13/viper"
	"reflect"
)

var conf *Config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Logger   Logger   `mapstructure:"logger"`
	Postgres Postgres `mapstructure:"postgres"`
}

type Server struct {
	Address string `mapstructure:"address"`
}

type Logger struct {
	Level    string `mapstructure:"level"`
	Filepath string `mapstructure:"filepath"`
	JSON     bool   `mapstructure:"json_form"`
	Stdout   bool   `mapstructure:"use_stdout"`
}

type Postgres struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"database"`
	MaxConns int    `mapstructure:"max_conns"`
}

func newConfig(confPath ...string) (c *Config) {
	viper.SetConfigName("api")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/pgram-back/")
	viper.AddConfigPath("./configs/")
	viper.AddConfigPath(".")

	for _, path := range confPath {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	c = &Config{}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err.Error())
	}

	return
}

func getAllStructTags(prefix, delim string, t reflect.Type) (res []string) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fieldName, ok := field.Tag.Lookup("mapstructure")
		if !ok {
			fieldName = field.Name
		}

		if field.Type.Kind() == reflect.Struct {
			res = append(res, getAllStructTags(fieldName, ".", field.Type)...)
			continue
		}

		res = append(res, prefix+delim+fieldName)
	}
	return
}

func checkUnsetFields(conf *Config) {
	for _, tag := range getAllStructTags("", "", reflect.TypeOf(*conf)) {
		if !viper.IsSet(tag) {
			panic("config field " + tag + "was not set")
		}
	}
}

func InitTestConfig(c *Config) {
	conf = c
}

// C provides config structure, can panic
func C() Config {
	if conf == nil {
		conf = newConfig()
		checkUnsetFields(conf)
	}

	return *conf
}
