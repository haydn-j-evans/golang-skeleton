package main

import (
	"flag"
	"fmt"

	featureflags "github.com/haydn-j-evans/go-skeleton/pkg/flags"
	"github.com/haydn-j-evans/go-skeleton/pkg/options"
)

const defaultOptionsFile string = ""

func main() {

	configFile := flag.String("config", defaultOptionsFile, "Define the location of the config yaml used in the application")
	flag.Parse()

	appOptions := options.InitOptions(configFile)

	flags := featureflags.InitFeatureFLags()

	featureflags.WatchFeatureFlags(flags)

	fmt.Println(*appOptions)
}
