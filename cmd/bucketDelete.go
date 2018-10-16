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
	"fmt"

	"github.com/spf13/cobra"
	"log"
	"os"
	"github.com/apprentice3d/forge-api-go-client/dm"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a bucket by providing bucket key",
	Long: `Delete a bucket, or several buckets by providing the bucket keys. 
For example:
	$forge-cli bucket delete some_bucket and_another_bucket

WARNING: This is an irreversible operation!
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("missing bucketKey as argument")
		}

		clientID := os.Getenv("FORGE_CLIENT_ID")
		clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
		bucketAPI := dm.NewBucketAPIWithCredentials(clientID, clientSecret)

		for _, bucketKey := range args {
			fmt.Println()
			fmt.Printf("Deleting bucket [%s]: ", bucketKey)

			err := bucketAPI.DeleteBucket(bucketKey)
			if err != nil {
				fmt.Println("FAIL:", err.Error())
				continue
			}
			fmt.Println("Done!")
		}

	},
}

func init() {
	bucketCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
