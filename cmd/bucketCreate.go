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
	"github.com/olekukonko/tablewriter"
	"time"
)


var policy string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("missing bucketKey as argument")
		}

		clientID := os.Getenv("FORGE_CLIENT_ID")
		clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
		bucketAPI := dm.NewBucketAPIWithCredentials(clientID, clientSecret)

		for _, bucketKey := range args {
			fmt.Println()
			fmt.Printf("Creating bucket [%s]: ", bucketKey)

			bucket, err := bucketAPI.CreateBucket(bucketKey, policy)
			if err != nil {
				fmt.Println("FAIL:", err.Error())
				continue
			}
			fmt.Println("Done")
			fmt.Println()
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Params", "Values"})
			utcDate := time.Unix(bucket.CreateDate, 0)

			data := [][]string{
				{"BucketKey", bucket.BucketKey},
				{"BucketOwner", bucket.BucketOwner},
				{"CreateDate", utcDate.String()},
				//{"Permissions", bucket.Permissions},
				{"PolicyKey", bucket.PolicyKey},
			}

			table.AppendBulk(data)
			table.SetBorder(false)
			table.Render()
		}
	},
}

func init() {
	bucketCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&policy,
		"policy",
		"p",
		"transient",
		"set bucket policy {transient, temporary, persistent}")

}
