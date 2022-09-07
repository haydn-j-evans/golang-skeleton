package options

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/drone/envsubst"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

type options struct {
	MetricsServer http.Server    `json:"metricsserver"`
	Server        http.Server    `json:"server"`
	Logging       loggingOptions `json:"loggingoptions"`
}

type loggingOptions struct {
	ReportCaller  bool                 `json:"reportcaller"`
	FormatterType string               `json:"formattertype"`
	LogLevel      string               `json:"loglevel"`
	JsonFormatter logrus.JSONFormatter `json:"jsonoptions"`
	TextFormatter logrus.TextFormatter `json:"textoptions"`
}

func InitOptions(configFile *string) *options {

	options := &options{}

	options.Server = http.Server{
		Addr: ":8080",
	}

	options.MetricsServer = http.Server{
		Addr: ":8081",
	}

	options.Logging.FormatterType = "textFormatter"

	options.Logging.ReportCaller = false

	options.Logging.LogLevel = "debug"

	options.Logging.TextFormatter = logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	}

	options.Logging.JsonFormatter = logrus.JSONFormatter{}

	err := LoadYaml(configFile, options)
	if err != nil {
		os.Exit(1)
	}

	return options
}

func LoadYaml(configFile *string, options *options) error {

	if *configFile == "" {
		fmt.Println("No configuration file provided! Using default values.")
		return nil
	}

	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Println("unable to load config file")
		return err
	}

	datastring, err := envsubst.EvalEnv(string(data))
	if err != nil {
		return fmt.Errorf("error in substituting env variables : %w", err)
	}

	data = []byte(datastring)
	// UnmarshalStrict will return an error if the config includes options that are
	// not mapped to felds of the into struct
	if err := yaml.UnmarshalStrict(data, options, yaml.DisallowUnknownFields); err != nil {
		return fmt.Errorf("error unmarshalling config: %w", err)
	}

	return nil

}
