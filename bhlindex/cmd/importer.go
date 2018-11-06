// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"os"

	"github.com/gnames/bhlindex/importer"
	"github.com/spf13/cobra"
)

// importerCmd represents the importer command
var importerCmd = &cobra.Command{
	Use:   "importer",
	Short: "Runs gRPC import server",
	Long: `Creates a gRPC import server on a specified port.
`,
	Run: func(cmd *cobra.Command, args []string) {
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		log.Printf("Starting gRPC server on %d port", port)

		importer.Serve(port)
	},
}

func init() {
	rootCmd.AddCommand(importerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	importerCmd.Flags().IntP("port", "p", 8889, "port for gRPC service")
}
