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
	"path/filepath"
	"github.com/apprentice3d/forge-api-go-client/dm"
	"io/ioutil"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "uploads an object into a bucket",
	Long: `upload expects 2 arguments [pathToFile] and [bucketId
For example:
	forge-cli bucket upload ./model.rvt sample-cache

will upload the model.rvt from current directory into sample-cache bucket.

`,
	Args:cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		clientID := os.Getenv("FORGE_CLIENT_ID")
		clientSecret := os.Getenv("FORGE_CLIENT_SECRET")
		bucketAPI := dm.NewBucketAPIWithCredentials(clientID, clientSecret)

		fmt.Printf("Uploading objects into bucket [%s]:\n",bucket)

		for _, filePath := range args[1:] {
			_,filename := filepath.Split(filePath)
			fmt.Printf("Uploading file %s: ", filename)
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				continue
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				continue
			}
			result, err := bucketAPI.UploadObject(bucket, filename, data)
			if err != nil {
				fmt.Println("ERROR:", err.Error())
				continue
			}

			fmt.Printf("Done uploading file with size %d Kb\n", result.Size/1024)
		}
	},
}

func init() {
	bucketCmd.AddCommand(uploadCmd)
}
