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
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	tags "github.com/wtolson/go-taglib"
)

//Audio holds the metadata values we're looking for
type Audio struct {
	Artist string
	Album  string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List metadata associated with the files in a directory.",
	Run:   runCmdList,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func runCmdList(cmd *cobra.Command, args []string) {
	res := listMetadata(directory)
	printResults(res)
}

func listMetadata(d string) []*Audio {
	files, err := ioutil.ReadDir(d)
	if err != nil {
		log.Printf("Issue reading directory: %v\n", err)
	}
	arr := []*Audio{}
	for _, f := range files {
		//Need to ensure the full file path is passed
		fs, err := tags.Read(d + "/" + f.Name())
		if err != nil {
			//If we get here, the file isn't something we care about so skip it
			continue

		}
		defer fs.Close()

		aa := Audio{fs.Artist(), fs.Album()}
		arr = append(arr, &aa)

	}

	return arr

}

func printResults(r []*Audio) {
	for i := range r {
		met := r[i]
		fmt.Printf("Artist: %v,  Album: %v\n", met.Artist, met.Album)
	}
}
