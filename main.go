package main

import (
	"flag"
	"fmt"

	"github.com/haydn-j-evans/go-skeleton/pkg/options"
)

const defaultOptionsFile string = "./config/options.yaml"

func main() {

	configFile := flag.String("config", defaultOptionsFile, "Define the location of the config yaml used in the application")
	flag.Parse()

	appOptions := options.InitOptions(configFile)

	appOptions.FeatureFlags.mutex.Rlock
	fmt.Println(*appOptions)
}
