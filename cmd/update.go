// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"io/ioutil"

	"fmt"

	"github.com/spf13/cobra"
	tags "github.com/wtolson/go-taglib"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update metadata using the global flags.",
	Run: func(cmd *cobra.Command, args []string) {
		files, _ := ioutil.ReadDir("./")
		for _, f := range files {
			fs, err := tags.Read(f.Name())
			if err != nil {
				//log.Printf("Can't convert %v: %v", f.Name(), err)
				continue

			}
			if artist != "" {
				fmt.Printf("Setting Artist %v\n", f.Name())
				fs.SetArtist(artist)
			}

			if album != "" {
				fmt.Printf("Setting Album %v\n", f.Name())

				fs.SetAlbum(album)
			}
			fs.Save()
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
