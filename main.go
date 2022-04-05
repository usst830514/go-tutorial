package main

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func readingConfigFiles() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error read config file: %w \n", err))
	}
}

func watching() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func writingConfigFiles() {
	if err := viper.SafeWriteConfigAs("/tmp/config.yaml"); err != nil {
		panic(fmt.Errorf("Fatal error write config file: %w \n", err))
	}
}

func readingConfigFromIOReader() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	if err := viper.ReadConfig(bytes.NewBuffer(yamlExample)); err != nil {
		panic(fmt.Errorf("Fatal error read config file: %w \n", err))
	}

	fmt.Println(viper.Get("hobbies")) // this would be "steve"
}

func registerAlias() {
	viper.RegisterAlias("loud", "Verbose")

	viper.Set("verbose", true) // same result as next line
	viper.Set("loud", true)    // same result as prior line

	viper.GetBool("loud")    // true
	viper.GetBool("verbose") // true
}

func main() {
	//readingConfigFiles()

	//watching()

	//writingConfigFiles()

	readingConfigFromIOReader()

}
