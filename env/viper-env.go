package env

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ViperEnv struct {
	Env *viper.Viper
}

var instance *ViperEnv

func NewEnv() *ViperEnv {
	if instance != nil {
		return instance
	}
	instance = &ViperEnv{Env: viper.New()}
	return instance
}

func (v *ViperEnv) Init() error {
	pflag.String("mode", "dev", "Microservice run mode")
	pflag.Parse()
	viper.AddConfigPath("")
	v.Env.BindPFlags(pflag.CommandLine)

	envMode := v.Env.GetString("mode")

	var configFile string = fmt.Sprintf("config.%v", envMode)

	v.Env.SetConfigName(configFile)
	v.Env.SetConfigType("json")
	v.Env.AddConfigPath("./config")

	if err := v.Env.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found", configFile)
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config file was found but another error was produced")
		}
		return err
	}
	fmt.Printf("%s config file loaded successfully", v.Env.Get("ENV"))
	fmt.Println()
	return nil
}

func (v *ViperEnv) Get(key string) interface{} {
	return v.Env.Get(key)
}
