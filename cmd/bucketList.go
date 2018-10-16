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
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"log"
	"github.com/apprentice3d/forge-api-go-client/dm"
	"github.com/olekukonko/tablewriter"
	"time"
	"strconv"
)

var shortFlag bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all buckets associated with current Forge secrets",
	Long: `list (forge-cli bucket list) will show all buckets associated with current Forge secrets. 

For example:
	$forge-cli bucket list
`,
	Run: func(cmd *cobra.Command, args []string) {
		clientID := os.Getenv("FORGE_CLIENT_ID")
		clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
		bucketAPI := dm.NewBucketAPIWithCredentials(clientID, clientSecret)
		list, err := bucketAPI.ListBuckets("US", "100", "")
		if err != nil {
			log.Fatal("Could not list the buckets: ", err.Error())
		}

		if shortFlag {

			for _, bucket := range list.Items {
				fmt.Println(bucket.BucketKey)
			}
			return
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"No.", "Bucket Key", "Create Date", "Policy"})
		var data [][]string
		for idx, bucket := range list.Items {
			utcDate := time.Unix(bucket.CreatedDate/1000, 0)
			data = append(data, []string{strconv.Itoa(idx + 1), bucket.BucketKey, utcDate.String(), bucket.PolicyKey})
		}
		table.AppendBulk(data)
		fmt.Println()
		fmt.Println("List of buckets associated with", clientID)
		table.Render()
	},
}

func init() {
	bucketCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&shortFlag,
		"short",
		"s",
		false,
		"display just the list with bucket keys")
}
