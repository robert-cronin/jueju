// Copyright 2024 Robert Cronin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config // viper is a popular configuration library for Go

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitConfig() {
	// Load base configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading base config file, %s", err)
	}

	// Get the environment variable
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Load environment-specific configuration
	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("Error reading %s config file, %s", env, err)
	}
}
