// Copyright (c) 2024 space-code
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
	echoserver "github.com/space-code/go-auth/internal/pkg/echo/server"
	"github.com/space-code/go-auth/internal/pkg/utils"
	"github.com/spf13/viper"
)

// Config holds configuration settings for the service.
type Config struct {
	// Name of the service.
	ServiceName string `mapstructure:"serviceName"`
	// Configuration for the Echo server.
	Echo *echoserver.EchoConfig `mapstructure:"echo"`
}

// configPath holds the path to the microservices configuration file.
var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "products write microservices config path")
}

// InitConfig initializes and returns a new configuration by reading from a file.
func InitConfig() (*Config, *echoserver.EchoConfig, error) {
	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "development"
	}

	if configPath == "" {
		configPathFromEnv := os.Getenv("GONFIG_PATH")

		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			d, err := utils.Dirname()

			if err != nil {
				return nil, nil, err
			}

			configPath = d
		}
	}

	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, cfg.Echo, nil
}
