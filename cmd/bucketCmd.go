// Copyright © 2018 Autodesk <denis.grigor@autodesk.com>
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

	"github.com/spf13/cobra"
)

// bucketCmd represents the bucket command
var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "bucket sub-command allows interaction with bucket manipulation part of the Data Management Service",
	Long: `bucket (forge-cli bucket) is a sub-command allowing interaction 
with bucket manipulation part of the Data Management Service. 

For example:
	forge-cli bucket list -> results in listing all buckets associated with current Forge secrets.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("bucket called")
	//},
}

func init() {
	rootCmd.AddCommand(bucketCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bucketCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bucketCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
