package zzz_config

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config     string
	configpath string
)

func init() {
	logrus.Info("starting zzz_config middleware init")

	flag.StringVar(&config, "config", "config", "config file name")
	flag.StringVar(&configpath, "configpath", ".", "the location of your config file")
	flag.Parse()

	//setup config
	viper.AddConfigPath(configpath)
	viper.AddConfigPath("/go/src")
	viper.SetConfigName(config)
	viper.ReadInConfig()

}
