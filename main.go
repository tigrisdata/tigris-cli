// Copyright 2022 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/tigrisdata/tigrisdb-cli/client"
	"github.com/tigrisdata/tigrisdb-cli/cmd"
	"github.com/tigrisdata/tigrisdb-cli/config"
	"github.com/tigrisdata/tigrisdb-cli/util"
)

func main() {
	util.LogConfigure()

	config.Load("tigris", &config.DefaultConfig)

	util.DefaultTimeout = config.DefaultConfig.Timeout

	if len(os.Args) > 2 && (os.Args[1] == "db" && os.Args[2] != "local") {
		if err := client.Init(config.DefaultConfig); err != nil {
			util.Error(err, "tigrisdb client initialization failed")
		}
	}

	cmd.Execute()
}