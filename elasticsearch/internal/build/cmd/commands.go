// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/zhangdapeng520/zdpgo_elasticsearch/elasticsearch/internal/build/utils"
)

var rootCmd = &cobra.Command{
	Use:   "build",
	Short: "Build tools for Elasticsearch client, helpers for CI, etc...",
}

// Execute launches the CLI application.
//
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PrintErr(err)
		os.Exit(1)
	}
}

// RegisterCmd adds a command to rootCmd.
//
func RegisterCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}