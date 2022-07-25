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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tigrisdata/tigris-cli/client"
	"github.com/tigrisdata/tigris-cli/util"
	"github.com/tigrisdata/tigris-client-go/driver"
)

var subscribeCmd = &cobra.Command{
	Use:     "subscribe {db} {collection}",
	Short:   "Subscribes to published messages",
	Long:    "Streams messages in real-time until cancelled.",
	Example: fmt.Sprintf("%[1]s subscribe testdb", rootCmd.Root().Name()),
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		it, err := client.Get().UseDatabase(args[0]).Subscribe(ctx, args[1])
		if err != nil {
			util.Error(err, "subscribe messages failed")
		}
		var doc driver.Document
		for it.Next(&doc) {
			util.Stdout("%s\n", string(doc))
		}
		if err := it.Err(); err != nil {
			util.Error(err, "iterate messages failed")
		}
	},
}

func init() {
	dbCmd.AddCommand(subscribeCmd)
}
