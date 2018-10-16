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
	"os"
	"log"
	"strconv"

	"github.com/apprentice3d/forge-api-go-client/dm"
	"github.com/olekukonko/tablewriter"
	"encoding/base64"
)

var encoded bool

// contentCmd represents the content command
var contentCmd = &cobra.Command{
	Use:   "content",
	Short: "lists the object contained in a given bucket",
	Long: `content will show all objects contained in a bucket, given the bucketKey. 

For example:
	$forge-cli bucket content my_experiments

will return the object contained in <my_experiments> bucket 
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			log.Fatal("missing bucketKey as argument")
		}

		clientID := os.Getenv("FORGE_CLIENT_ID")
		clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
		bucketAPI := dm.NewBucketAPIWithCredentials(clientID, clientSecret)

		for _,bucketKey := range args {
			fmt.Println()
			fmt.Printf("Listing content for [%s] bucket:\n", bucketKey)
			list, err := bucketAPI.ListObjects(bucketKey, "", "", "")
			if err != nil {
				log.Fatal("Could not list the bucket content: ", err.Error())
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"No.", "ObjectKey", "ObjectID", "Size"})
			var data [][]string
			var totalSize uint64
			for idx, object := range list.Items {
				totalSize += object.Size
				objectID := object.ObjectID
				if encoded { objectID = base64.RawStdEncoding.EncodeToString([]byte(objectID))}
				data = append(data, []string{strconv.Itoa(idx + 1),
					object.ObjectKey,
					objectID,
					strconv.FormatUint(object.Size, 10),
				})
			}


			table.SetFooter([]string{"","","Total", fmt.Sprintf("%.2f Mb", float64(totalSize)/(1024*1024))})
			table.AppendBulk(data)
			table.Render()
		}



	},
}

func init() {
	bucketCmd.AddCommand(contentCmd)


	contentCmd.Flags().BoolVarP(&encoded,
		"encoded",
		"e",
		false,
		"show the ObjectID as Base64-encoded URN")
}
