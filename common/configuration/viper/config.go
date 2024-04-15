package viper

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var VPMap map[string]*viper.Viper

type ConfigStruct struct {
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
	Type string `mapstructure:"type"`
}

type Configuration struct {
	Configs []ConfigStruct `mapstructure:"configs"`
}

func LoadMainConfig() map[string]*viper.Viper {
	VPMap = make(map[string]*viper.Viper)
	vp := viper.New()
	vp.SetConfigName("configuration") // name of config file (without extension)
	vp.SetConfigType("json")          // REQUIRED if the config file does not have the extension in the name
	vp.AddConfigPath(".")             // optionally look for config in the working directory
	err := vp.ReadInConfig()          // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("fatal error - No configuration file: %w", err))
	}
	VPMap["configuration"] = vp
	VPMap = loadAllConfig(vp)

	return VPMap
}

func GetViper(name string) *viper.Viper {
	return VPMap[name]
}

func loadAllConfig(vp *viper.Viper) map[string]*viper.Viper {
	var configuration Configuration
	err := vp.Unmarshal(&configuration)
	if err == nil && configuration.Configs != nil {
		for i, c := range configuration.Configs {
			if vp := createViper(c); vp != nil {
				VPMap[c.Name] = vp
				fmt.Println(i, c)
			}
		}
	}

	return VPMap
}

func createViper(configStruct ConfigStruct) *viper.Viper {
	var vp *viper.Viper

	vp = viper.New()
	if configStruct.Name != "" {
		vp.SetConfigName(configStruct.Name)
	}
	if configStruct.Type != "" {
		vp.SetConfigType(configStruct.Type)
	}
	if configStruct.Path != "" {
		vp.AddConfigPath(configStruct.Path)
	}
	err := vp.ReadInConfig() // Find and read the config file
	if err != nil {          // Handle errors reading the config file
		panic(fmt.Errorf("fatal error - No configuration file: %w", err))
	}

	return vp
}

func addKey(vp *viper.Viper, key string, value string) {
	vp.Set(key, value)
	vp.WriteConfig()
}

func addWatchConfig(vp *viper.Viper) {
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("file changed: %s\n", in.Name)
	})
	vp.WatchConfig()
}
